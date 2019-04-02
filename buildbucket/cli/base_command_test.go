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

package cli

import (
	"fmt"
	"testing"

	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestRetrieveBuildIDs(t *testing.T) {
	t.Parallel()

	Convey("RetrieveBuildIDs", t, func() {
		Convey("Basic", func() {
			builds := []string{
				"1",
				"a/b/c/2",
				"a/b/c/3",
			}
			var actualReq *pb.BatchRequest
			ids, err := retrieveBuildIDs(builds, func(req *pb.BatchRequest) (*pb.BatchResponse, error) {
				actualReq = req
				return textpb(&pb.BatchResponse{}, `
					responses { get_build {id: 2} }
					responses { get_build {id: 3} }
				`).(*pb.BatchResponse), nil
			})
			So(err, ShouldBeNil)
			So(actualReq, ShouldResembleProtoText, `
				requests {
					get_build {
						builder { project: "a" bucket: "b" builder: "c"}
						build_number: 2
						fields { paths: "id" }
					}
				}
				requests {
					get_build {
						builder { project: "a" bucket: "b" builder: "c"}
						build_number: 3
						fields { paths: "id" }
					}
				}
			`)
			So(ids, ShouldResemble, []int64{1, 2, 3})
		})

		Convey("No build numbers", func() {
			builds := []string{"1", "2"}
			var actualReq *pb.BatchRequest
			ids, err := retrieveBuildIDs(builds, func(req *pb.BatchRequest) (*pb.BatchResponse, error) {
				actualReq = req
				return nil, fmt.Errorf("unexpected")
			})
			So(err, ShouldBeNil)
			So(actualReq, ShouldBeNil)
			So(ids, ShouldResemble, []int64{1, 2})
		})
	})
}