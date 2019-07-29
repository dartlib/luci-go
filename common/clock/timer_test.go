// Copyright 2015 The LUCI Authors.
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

package clock

import (
	"context"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimerResult(t *testing.T) {
	t.Parallel()

	Convey(`Testing TimerResult`, t, func() {
		Convey(`A TimerResult with no error is not incomplete.`, func() {
			So(TimerResult{}.Incomplete(), ShouldBeFalse)
		})

		Convey(`A TimerResult with context.Canceled, context.DeadlineExceeded, or another error is incomplete.`, func() {
			So(TimerResult{Err: context.Canceled}.Incomplete(), ShouldBeTrue)
			So(TimerResult{Err: context.DeadlineExceeded}.Incomplete(), ShouldBeTrue)
			So(TimerResult{Err: errors.New("foo")}.Incomplete(), ShouldBeTrue)
		})
	})
}
