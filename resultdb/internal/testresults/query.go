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

package testresults

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/spanner"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/trace"

	"go.chromium.org/luci/resultdb/internal/invocations"
	"go.chromium.org/luci/resultdb/internal/pagination"
	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/pbutil"
	pb "go.chromium.org/luci/resultdb/proto/v1"
)

// Query specifies test results to fetch.
type Query struct {
	InvocationIDs     invocations.IDSet
	Predicate         *pb.TestResultPredicate
	PageSize          int // must be positive
	PageToken         string
	SelectVariantHash bool
}

func (q *Query) run(ctx context.Context, txn *spanner.ReadOnlyTransaction, f func(QueryItem) error) (err error) {
	ctx, ts := trace.StartSpan(ctx, "QueryTestResults.run")
	defer func() { ts.End(err) }()

	if q.PageSize < 0 {
		panic("PageSize < 0")
	}

	var extraSelect []string
	if q.SelectVariantHash {
		extraSelect = append(extraSelect, "tr.VariantHash")
	}

	from := "TestResults tr"
	if q.Predicate.GetExpectancy() == pb.TestResultPredicate_VARIANTS_WITH_UNEXPECTED_RESULTS {
		// We must return only test results of test variants that have unexpected results.
		//
		// The following query ensures that we first select test variants with
		// unexpected results, and then for each variant do a lookup in TestResults
		// table.
		from = `
			VariantsWithUnexpectedResults vur
			JOIN@{FORCE_JOIN_ORDER=TRUE, JOIN_METHOD=HASH_JOIN} TestResults tr USING (TestId, VariantHash)
			`
	}

	limit := ""
	if q.PageSize > 0 {
		limit = `LIMIT @limit`
	}

	st := spanner.NewStatement(fmt.Sprintf(`
		@{USE_ADDITIONAL_PARALLELISM=TRUE}
		WITH VariantsWithUnexpectedResults AS (
			# Note: this query is not executed if it ends up not used in the top-level
			# query.
			SELECT DISTINCT TestId, VariantHash
			FROM TestResults@{FORCE_INDEX=UnexpectedTestResults}
			WHERE IsUnexpected AND InvocationId IN UNNEST(@invIDs)
		)
		SELECT
			tr.InvocationId,
			tr.TestId,
			tr.ResultId,
			tr.Variant,
			tr.IsUnexpected,
			tr.Status,
			tr.SummaryHtml,
			tr.StartTime,
			tr.RunDurationUsec,
			tr.Tags,
			%s
		FROM %s
		WHERE InvocationId IN UNNEST(@invIDs)
			# Skip test results after the one specified in the page token.
			AND (
				(tr.InvocationId > @afterInvocationId) OR
				(tr.InvocationId = @afterInvocationId AND tr.TestId > @afterTestId) OR
				(tr.InvocationId = @afterInvocationId AND tr.TestId = @afterTestId AND tr.ResultId > @afterResultId)
			)
			AND REGEXP_CONTAINS(tr.TestId, @TestIdRegexp)
			AND (@variantHashEquals IS NULL OR tr.VariantHash = @variantHashEquals)
			AND (@variantContains IS NULL
				OR ARRAY_LENGTH(@variantContains) = 0
				OR (SELECT LOGICAL_AND(kv IN UNNEST(tr.Variant)) FROM UNNEST(@variantContains) kv)
			)
		ORDER BY InvocationId, TestId, ResultId
		%s
	`, strings.Join(extraSelect, ","), from, limit))
	st.Params["invIDs"] = q.InvocationIDs
	st.Params["limit"] = q.PageSize

	// Filter by test id.
	testIDRegexp := q.Predicate.GetTestIdRegexp()
	if testIDRegexp == "" {
		testIDRegexp = ".*"
	}
	st.Params["TestIdRegexp"] = fmt.Sprintf("^%s$", testIDRegexp)

	// Filter by variant.
	PopulateVariantParams(&st, q.Predicate.GetVariant())

	// Apply page token.
	err = invocations.TokenToMap(q.PageToken, st.Params, "afterInvocationId", "afterTestId", "afterResultId")
	if err != nil {
		return err
	}

	// Read the results.
	var summaryHTML span.Compressed
	var b span.Buffer
	return span.Query(ctx, txn, st, func(row *spanner.Row) error {
		var invID invocations.ID
		var maybeUnexpected spanner.NullBool
		var micros spanner.NullInt64
		tr := &pb.TestResult{}
		item := QueryItem{TestResult: tr}

		ptrs := []interface{}{
			&invID,
			&tr.TestId,
			&tr.ResultId,
			&tr.Variant,
			&maybeUnexpected,
			&tr.Status,
			&summaryHTML,
			&tr.StartTime,
			&micros,
			&tr.Tags,
		}
		if q.SelectVariantHash {
			ptrs = append(ptrs, &item.VariantHash)
		}

		err = b.FromSpanner(row, ptrs...)
		if err != nil {
			return err
		}

		tr.Name = pbutil.TestResultName(string(invID), tr.TestId, tr.ResultId)
		tr.SummaryHtml = string(summaryHTML)
		populateExpectedField(tr, maybeUnexpected)
		populateDurationField(tr, micros)

		return f(item)
	})
}

// Fetch returns a page of test results matching q.
// Returned test results are ordered by parent invocation ID, test ID and result
// ID.
func (q *Query) Fetch(ctx context.Context, txn *spanner.ReadOnlyTransaction) (trs []*pb.TestResult, nextPageToken string, err error) {
	if q.PageSize <= 0 {
		panic("PageSize <= 0")
	}

	trs = make([]*pb.TestResult, 0, q.PageSize)
	err = q.run(ctx, txn, func(item QueryItem) error {
		trs = append(trs, item.TestResult)
		return nil
	})
	if err != nil {
		trs = nil
		return
	}

	// If we got pageSize results, then we haven't exhausted the collection and
	// need to return the next page token.
	if len(trs) == q.PageSize {
		last := trs[q.PageSize-1]
		invID, testID, resultID := MustParseName(last.Name)
		nextPageToken = pagination.Token(string(invID), testID, resultID)
	}
	return
}

// QueryItem is one element returned by a Query.
type QueryItem struct {
	*pb.TestResult
	VariantHash string
}

// Run calls f for test results matching the query.
// The test results are ordered by parent invocation ID, test ID and result ID.
func (q *Query) Run(ctx context.Context, txn *spanner.ReadOnlyTransaction, f func(QueryItem) error) error {
	if q.PageSize > 0 {
		panic("PageSize is specified when Query.Run")
	}
	return q.run(ctx, txn, f)
}

// PopulateVariantParams populates variantHashEquals and variantContains
// parameters based on the predicate.
func PopulateVariantParams(st *spanner.Statement, variantPredicate *pb.VariantPredicate) {
	st.Params["variantHashEquals"] = spanner.NullString{}
	st.Params["variantContains"] = []string(nil)
	switch p := variantPredicate.GetPredicate().(type) {
	case *pb.VariantPredicate_Equals:
		st.Params["variantHashEquals"] = pbutil.VariantHash(p.Equals)
	case *pb.VariantPredicate_Contains:
		st.Params["variantContains"] = pbutil.VariantToStrings(p.Contains)
	case nil:
		// No filter.
	default:
		panic(errors.Reason("unexpected variant predicate %q", variantPredicate).Err())
	}
}
