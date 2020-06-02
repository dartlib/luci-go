// Copyright 2020 The LUCI Authors.
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

package exonerations

import (
	"context"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/grpc/appstatus"

	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"
)

// MustParseName extracts invocation, test id and exoneration
// IDs from the name.
// Panics on failure.
func MustParseName(name string) (invID span.InvocationID, testID, exonerationID string) {
	invIDStr, testID, exonerationID, err := pbutil.ParseTestExonerationName(name)
	if err != nil {
		panic(err)
	}
	invID = span.InvocationID(invIDStr)
	return
}

// Read reads a test exoneration from Spanner.
// If it does not exist, the returned error is annotated with NotFound GRPC
// code.
func Read(ctx context.Context, txn span.Txn, name string) (*pb.TestExoneration, error) {
	invIDStr, testID, exonerationID, err := pbutil.ParseTestExonerationName(name)
	if err != nil {
		return nil, err
	}
	invID := span.InvocationID(invIDStr)

	ret := &pb.TestExoneration{
		Name:          name,
		TestId:        testID,
		ExonerationId: exonerationID,
	}

	// Populate fields from TestExonerations table.
	var explanationHTML span.Compressed
	err = span.ReadRow(ctx, txn, "TestExonerations", invID.Key(testID, exonerationID), map[string]interface{}{
		"Variant":         &ret.Variant,
		"ExplanationHTML": &explanationHTML,
	})
	switch {
	case spanner.ErrCode(err) == codes.NotFound:
		return nil, appstatus.Attachf(err, codes.NotFound, "%s not found", ret.Name)

	case err != nil:
		return nil, errors.Annotate(err, "failed to fetch %q", ret.Name).Err()

	default:
		ret.ExplanationHtml = string(explanationHTML)
		return ret, nil
	}
}
