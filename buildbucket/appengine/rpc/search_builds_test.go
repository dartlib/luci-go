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

package rpc

import (
	"testing"

	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestSearchBuilds(t *testing.T) {
	t.Parallel()

	Convey("validateChange", t, func() {
		Convey("nil", func() {
			err := validateChange(nil)
			So(err, ShouldErrLike, "host is required")
		})

		Convey("empty", func() {
			ch := &pb.GerritChange{}
			err := validateChange(ch)
			So(err, ShouldErrLike, "host is required")
		})

		Convey("change", func() {
			ch := &pb.GerritChange{
				Host: "host",
			}
			err := validateChange(ch)
			So(err, ShouldErrLike, "change is required")
		})

		Convey("patchset", func() {
			ch := &pb.GerritChange{
				Host:   "host",
				Change: 1,
			}
			err := validateChange(ch)
			So(err, ShouldErrLike, "patchset is required")
		})

		Convey("valid", func() {
			ch := &pb.GerritChange{
				Host:     "host",
				Change:   1,
				Patchset: 1,
			}
			err := validateChange(ch)
			So(err, ShouldBeNil)
		})
	})

	Convey("validateCommit", t, func() {
		Convey("nil", func() {
			err := validateCommit(nil)
			So(err, ShouldErrLike, "host is required")
		})

		Convey("empty", func() {
			cm := &pb.GitilesCommit{}
			err := validateCommit(cm)
			So(err, ShouldErrLike, "host is required")
		})

		Convey("project", func() {
			cm := &pb.GitilesCommit{
				Host: "host",
			}
			err := validateCommit(cm)
			So(err, ShouldErrLike, "project is required")
		})

		Convey("id", func() {
			cm := &pb.GitilesCommit{
				Host:    "host",
				Project: "project",
				Id:      "id",
			}
			err := validateCommit(cm)
			So(err, ShouldErrLike, "id must match")
		})

		Convey("ref", func() {
			cm := &pb.GitilesCommit{
				Host:    "host",
				Project: "project",
				Ref:     "ref",
			}
			err := validateCommit(cm)
			So(err, ShouldErrLike, "ref must match")
		})

		Convey("mutual exclusion", func() {
			Convey("ref", func() {
				cm := &pb.GitilesCommit{
					Host:    "host",
					Project: "project",
					Id:      "id",
					Ref:     "ref",
				}
				err := validateCommit(cm)
				So(err, ShouldErrLike, "id is mutually exclusive with (ref and position)")
			})

			Convey("position", func() {
				cm := &pb.GitilesCommit{
					Host:     "host",
					Project:  "project",
					Id:       "id",
					Position: 1,
				}
				err := validateCommit(cm)
				So(err, ShouldErrLike, "id is mutually exclusive with (ref and position)")
			})

			Convey("neither", func() {
				cm := &pb.GitilesCommit{
					Host:    "host",
					Project: "project",
				}
				err := validateCommit(cm)
				So(err, ShouldErrLike, "one of")
			})
		})

		Convey("valid", func() {
			Convey("id", func() {
				cm := &pb.GitilesCommit{
					Host:    "host",
					Project: "project",
					Id:      "1234567890123456789012345678901234567890",
				}
				err := validateCommit(cm)
				So(err, ShouldBeNil)
			})

			Convey("ref", func() {
				cm := &pb.GitilesCommit{
					Host:     "host",
					Project:  "project",
					Ref:      "refs/ref",
					Position: 1,
				}
				err := validateCommit(cm)
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("validatePredicate", t, func() {
		Convey("nil", func() {
			err := validatePredicate(nil)
			So(err, ShouldBeNil)
		})

		Convey("empty", func() {
			pr := &pb.BuildPredicate{}
			err := validatePredicate(pr)
			So(err, ShouldBeNil)
		})

		Convey("mutual exclusion", func() {
			pr := &pb.BuildPredicate{
				Build:      &pb.BuildRange{},
				CreateTime: &pb.TimeRange{},
			}
			err := validatePredicate(pr)
			So(err, ShouldErrLike, "build is mutually exclusive with create_time")
		})
	})

	Convey("validateSearch", t, func() {
		Convey("nil", func() {
			err := validateSearch(nil)
			So(err, ShouldBeNil)
		})

		Convey("empty", func() {
			req := &pb.SearchBuildsRequest{}
			err := validateSearch(req)
			So(err, ShouldBeNil)
		})

		Convey("page size", func() {
			Convey("negative", func() {
				req := &pb.SearchBuildsRequest{
					PageSize: -1,
				}
				err := validateSearch(req)
				So(err, ShouldErrLike, "page_size cannot be negative")
			})

			Convey("zero", func() {
				req := &pb.SearchBuildsRequest{
					PageSize: 0,
				}
				err := validateSearch(req)
				So(err, ShouldBeNil)
			})

			Convey("positive", func() {
				req := &pb.SearchBuildsRequest{
					PageSize: 1,
				}
				err := validateSearch(req)
				So(err, ShouldBeNil)
			})
		})
	})
}
