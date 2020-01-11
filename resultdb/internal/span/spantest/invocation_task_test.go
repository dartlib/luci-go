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

package spantest

import (
	"testing"
	"time"

	"go.chromium.org/luci/common/clock"

	internalpb "go.chromium.org/luci/resultdb/internal/proto"
	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/internal/testutil"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSampleInvocationTasks(t *testing.T) {
	Convey(`TestSampleInvocationTasks`, t, func() {
		ctx := testutil.SpannerTestContext(t)
		now := clock.Now(ctx)
		invTask := &internalpb.InvocationTask{}

		testutil.MustApply(ctx,
			span.InsertInvocationTask(span.InvocationID("inv0"), "task_1", invTask, now.Add(-time.Hour), false),
			span.InsertInvocationTask(span.InvocationID("inv0"), "task_2", invTask, now.Add(-time.Hour), false),
			span.InsertInvocationTask(span.InvocationID("inv1"), "task_3", invTask, now.Add(-time.Hour), false),
			span.InsertInvocationTask(span.InvocationID("inv1"), "task_4", invTask, now, false),
			span.InsertInvocationTask(span.InvocationID("inv2"), "task_5", invTask, now.Add(time.Hour), false),
		)

		txn := span.Client(ctx).ReadOnlyTransaction()
		defer txn.Close()

		rows, err := span.SampleInvocationTasks(ctx, txn, now, 3)
		So(err, ShouldBeNil)
		So(rows, ShouldHaveLength, 3)
		So(rows, ShouldNotContain, span.TaskKey{
			InvocationID: span.InvocationID("inv0"),
			TaskID:       "task_5",
		})
	})
}