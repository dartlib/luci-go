// Copyright 2019 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deriver

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"cloud.google.com/go/spanner"
	durpb "github.com/golang/protobuf/ptypes/duration"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"

	swarmingAPI "go.chromium.org/luci/common/api/swarming/swarming/v1"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/grpc/appstatus"

	"go.chromium.org/luci/resultdb/internal"
	"go.chromium.org/luci/resultdb/internal/artifacts"
	"go.chromium.org/luci/resultdb/internal/invocations"
	"go.chromium.org/luci/resultdb/internal/services/deriver/chromium"
	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/internal/tasks"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/v1"
)

// testResultBatchSizeMax is the maximum number of TestResults to include per transaction.
// Note that the same transaction is used for both test results and artifacts.
const testResultBatchSizeMax = 500

var urlPrefixes = []string{"http://", "https://"}

// validateDeriveChromiumInvocationRequest returns an error if req is invalid.
func validateDeriveChromiumInvocationRequest(req *pb.DeriveChromiumInvocationRequest) error {
	if req.SwarmingTask == nil {
		return errors.Reason("swarming_task missing").Err()
	}

	if req.SwarmingTask.Hostname == "" {
		return errors.Reason("swarming_task.hostname missing").Err()
	}

	for _, prefix := range urlPrefixes {
		if strings.HasPrefix(req.SwarmingTask.Hostname, prefix) {
			return errors.Reason("swarming_task.hostname should not have prefix %q", prefix).Err()
		}
	}

	if req.SwarmingTask.Id == "" {
		return errors.Reason("swarming_task.id missing").Err()
	}

	return nil
}

// DeriveChromiumInvocation derives the invocation associated with the given swarming task.
//
// The invocation returned is associated with the swarming task itself.
// If the task is deduped against another task, the invocation returned includes the underlying one.
func (s *deriverServer) DeriveChromiumInvocation(ctx context.Context, in *pb.DeriveChromiumInvocationRequest) (*pb.Invocation, error) {
	if err := validateDeriveChromiumInvocationRequest(in); err != nil {
		return nil, appstatus.BadRequest(err)
	}

	// Get the swarming service to use.
	swarmingURL := "https://" + in.SwarmingTask.Hostname
	swarmSvc, err := chromium.GetSwarmSvc(internal.HTTPClient(ctx), swarmingURL)
	if err != nil {
		return nil, errors.Annotate(err, "creating swarming client for %q", swarmingURL).Err()
	}

	// Get the swarming task.
	task, err := chromium.GetSwarmingTask(ctx, in.SwarmingTask.Id, swarmSvc)
	if err != nil {
		return nil, errors.Annotate(err, "getting swarming task %q on %q",
			in.SwarmingTask.Id, in.SwarmingTask.Hostname).Err()
	}
	invID := chromium.GetInvocationID(task, in)

	client := span.Client(ctx)

	// Check if we need to write this invocation.
	switch err := shouldWriteInvocation(ctx, client.Single(), invID); {
	case err == errAlreadyExists:
		readTxn := client.ReadOnlyTransaction()
		defer readTxn.Close()
		return invocations.Read(ctx, readTxn, invID)
	case err != nil:
		return nil, err
	}

	inv, err := chromium.DeriveChromiumInvocation(task, in)
	if err != nil {
		return nil, err
	}

	// Derive the origin invocation and results.
	var originInv *pb.Invocation
	switch originInv, err = s.deriveInvocationForOriginTask(ctx, in, task, swarmSvc, client); {
	case err != nil:
		return nil, err
	case inv.Name == originInv.Name: // origin task is the task itself, we're done.
		return originInv, nil
	}

	// Include originInv into inv.
	inv.IncludedInvocations = []string{originInv.Name}
	invMs := []*spanner.Mutation{
		span.InsertMap("Invocations", s.rowOfInvocation(ctx, inv, "", 0)),
		span.InsertMap("IncludedInvocations", map[string]interface{}{
			"InvocationId":         invID,
			"IncludedInvocationId": invocations.MustParseName(originInv.Name),
		}),
	}
	_, err = span.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		if err := shouldWriteInvocation(ctx, txn, invID); err != nil {
			return err
		}
		return txn.BufferWrite(invMs)
	})
	switch {
	case err == errAlreadyExists:
		// Pass.
	case err != nil:
		return nil, err
	default:
		span.IncRowCount(ctx, 1, span.Invocations, span.Inserted)
	}

	return inv, nil
}

// errAlreadyExists is returned by shouldWriteInvocation if the invocation
// already exists.
var errAlreadyExists = fmt.Errorf("already exists")

// shouldWriteInvocation returns errAlreadyExists if the invocation already
// exists and should not be re-written.
func shouldWriteInvocation(ctx context.Context, txn span.Txn, id invocations.ID) error {
	state, err := invocations.ReadState(ctx, txn, id)
	s, _ := appstatus.Get(err)
	switch {
	case s.Code() == codes.NotFound:
		// No such invocation found means we may have to write it, so proceed.
		return nil

	case err != nil:
		return err

	case state != pb.Invocation_FINALIZED:
		return errors.Reason(
			"attempting to derive an existing non-finalized invocation").Err()

	default:
		// The invocation exists and is finalized, so no need to write it.
		return errAlreadyExists
	}
}

// deriveInvocationForOriginTask derives an invocation and test results
// from a given task and returns derived origin invocation.
func (s *deriverServer) deriveInvocationForOriginTask(ctx context.Context, in *pb.DeriveChromiumInvocationRequest, task *swarmingAPI.SwarmingRpcsTaskResult, swarmSvc *swarmingAPI.Service, client *spanner.Client) (*pb.Invocation, error) {
	// Get the origin task that the task is deduped against. Or the task
	// itself if it's not deduped.
	originTask, err := chromium.GetOriginTask(ctx, task, swarmSvc)
	if err != nil {
		return nil, errors.Annotate(err, "getting origin for swarming task %q on %q",
			in.SwarmingTask.Id, in.SwarmingTask.Hostname).Err()
	}
	originInvID := chromium.GetInvocationID(originTask, in)

	// Check if we need to write origin invocation.
	switch err := shouldWriteInvocation(ctx, client.Single(), originInvID); {
	case err == errAlreadyExists:
		txn := client.ReadOnlyTransaction()
		defer txn.Close()
		return invocations.Read(ctx, txn, originInvID)
	case err != nil:
		return nil, err
	}
	originInv, err := chromium.DeriveChromiumInvocation(originTask, in)
	if err != nil {
		return nil, err
	}

	// Get the protos and prepare to write them to Spanner.
	logging.Infof(ctx, "Deriving task %q on %q", originTask.TaskId, in.SwarmingTask.Hostname)
	results, err := chromium.DeriveTestResults(ctx, originTask, in, originInv)
	if err != nil {
		return nil, errors.Annotate(err,
			"task %q on %q named %q", in.SwarmingTask.Id, in.SwarmingTask.Hostname, originTask.Name).Err()
	}
	// TODO(jchinlee): Validate invocation and results.

	// Write test results in batches concurrently, updating inv with the names of the invocations
	// that will be included.
	batchInvs, err := s.batchInsertTestResults(ctx, originInv, results, testResultBatchSizeMax)
	if err != nil {
		return nil, err
	}
	originInv.IncludedInvocations = batchInvs.Names()

	// Prepare mutations.
	ms := make([]*spanner.Mutation, 0, len(batchInvs)+4)
	ms = append(ms, span.InsertMap("Invocations", s.rowOfInvocation(ctx, originInv, "", 0)))
	for includedID := range batchInvs {
		ms = append(ms, span.InsertMap("IncludedInvocations", map[string]interface{}{
			"InvocationId":         originInvID,
			"IncludedInvocationId": includedID,
		}))
	}
	ms = append(ms, tasks.EnqueueBQExport(originInvID, s.InvBQTable, clock.Now(ctx).UTC()))

	_, err = span.ReadWriteTransaction(ctx, func(ctx context.Context, txn *spanner.ReadWriteTransaction) error {
		// Check origin invocation state again.
		if err := shouldWriteInvocation(ctx, txn, originInvID); err != nil {
			return err
		}
		return txn.BufferWrite(ms)
	})
	switch {
	case err == errAlreadyExists:
		// Pass.
	case err != nil:
		return nil, err
	default:
		span.IncRowCount(ctx, 1, span.Invocations, span.Inserted)
	}

	return originInv, nil
}

// batchInsertTestResults inserts the given TestResults in batches under container Invocations,
// returning container ids.
func (s *deriverServer) batchInsertTestResults(ctx context.Context, inv *pb.Invocation, trs []*chromium.TestResult, batchSize int) (invocations.IDSet, error) {
	batches := batchTestResults(trs, batchSize)
	includedInvs := make(invocations.IDSet, len(batches))

	invID := invocations.MustParseName(inv.Name)
	eg, ctx := errgroup.WithContext(ctx)
	client := span.Client(ctx)
	for i, batch := range batches {
		i := i
		batch := batch

		batchID := batchInvocationID(invID, i)
		includedInvs.Add(batchID)

		eg.Go(func() error {
			ms := make([]*spanner.Mutation, 0, len(batch)+1)

			// Convert the container Invocation in the batch.
			batchInv := &pb.Invocation{
				Name:         batchID.Name(),
				State:        pb.Invocation_FINALIZED,
				CreateTime:   inv.CreateTime,
				FinalizeTime: inv.FinalizeTime,
				Deadline:     inv.Deadline,
				Tags:         inv.Tags,
			}
			ms = append(ms, span.InsertOrUpdateMap(
				"Invocations", s.rowOfInvocation(ctx, batchInv, "", int64(len(batch)))),
			)

			// Convert the TestResults in the batch.
			for k, tr := range batch {
				tr.ResultId = strconv.Itoa(k)
				ms = append(ms, insertOrUpdateTestResult(batchID, tr.TestResult))
				for _, a := range tr.Artifacts {
					ms = append(ms, insertOrUpdateArtifact(batchID, tr.TestResult, a))
				}
			}

			if _, err := client.Apply(ctx, ms); err != nil {
				return err
			}

			span.IncRowCount(ctx, len(batch), span.TestResults, span.Inserted)
			span.IncRowCount(ctx, 1, span.Invocations, span.Inserted)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return includedInvs, nil
}

// batchInvocationID returns an InvocationID for the Invocation containing the referenced batch.
func batchInvocationID(invID invocations.ID, batchInd int) invocations.ID {
	return invocations.ID(fmt.Sprintf("%s-batch-%d", invID, batchInd))
}

// batchTestResults batches the given TestResults given the maximum batch size.
func batchTestResults(trs []*chromium.TestResult, batchSize int) [][]*chromium.TestResult {
	batches := make([][]*chromium.TestResult, 0, len(trs)/batchSize+1)
	for len(trs) > 0 {
		end := batchSize
		if end > len(trs) {
			end = len(trs)
		}

		batches = append(batches, trs[:end])
		trs = trs[end:]
	}

	return batches
}

func insertOrUpdateTestResult(invID invocations.ID, tr *pb.TestResult) *spanner.Mutation {
	trMap := map[string]interface{}{
		"InvocationId": invID,
		"TestId":       tr.TestId,
		"ResultId":     tr.ResultId,

		"Variant":     tr.Variant,
		"VariantHash": pbutil.VariantHash(tr.Variant),

		"CommitTimestamp": spanner.CommitTimestamp,

		"Status":          tr.Status,
		"SummaryHTML":     span.Compressed([]byte(tr.SummaryHtml)),
		"StartTime":       tr.StartTime,
		"RunDurationUsec": toMicros(tr.Duration),
		"Tags":            tr.Tags,
	}

	// Populate IsUnexpected /only/ if true, to keep the index thin.
	if !tr.Expected {
		trMap["IsUnexpected"] = true
	}

	return span.InsertOrUpdateMap("TestResults", trMap)
}

func insertOrUpdateArtifact(invID invocations.ID, tr *pb.TestResult, a *pb.Artifact) *spanner.Mutation {
	return span.InsertOrUpdateMap("Artifacts", map[string]interface{}{
		"InvocationId": invID,
		"ParentId":     artifacts.ParentID(tr.TestId, tr.ResultId),
		"ArtifactId":   a.ArtifactId,
		"ContentType":  a.ContentType,
		"Size":         a.SizeBytes,
		"IsolateURL":   a.FetchUrl,
	})
}

// toMicros converts a duration.Duration proto to microseconds.
func toMicros(d *durpb.Duration) int64 {
	if d == nil {
		return 0
	}
	return 1e6*d.Seconds + int64(1e-3*float64(d.Nanos))
}
