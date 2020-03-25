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

package model

import (
	"context"
	"testing"

	structpb "github.com/golang/protobuf/ptypes/struct"

	"go.chromium.org/gae/impl/memory"
	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/common/clock/testclock"

	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestModel(t *testing.T) {
	t.Parallel()

	Convey("Build", t, func() {
		ctx := memory.Use(context.Background())
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)

		Convey("read/write", func() {
			So(datastore.Put(ctx, &Build{
				ID: 1,
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
				},
				Project:    "project",
				BucketID:   "project/bucket",
				BuilderID:  "project/bucket/builder",
				CreateTime: testclock.TestRecentTimeUTC,
				Status:     pb.Status_SUCCESS,
			}), ShouldBeNil)

			b := &Build{
				ID: 1,
			}
			So(datastore.Get(ctx, b), ShouldBeNil)
			So(b, ShouldResemble, &Build{
				ID: 1,
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
				},
				Project:    "project",
				BucketID:   "project/bucket",
				BuilderID:  "project/bucket/builder",
				CreateTime: datastore.RoundTime(testclock.TestRecentTimeUTC),
				Status:     pb.Status_SUCCESS,
			})
		})

		Convey("ToProto", func() {
			b := &Build{
				ID: 1,
				Proto: pb.Build{
					Id: 1,
				},
				Tags: []string{"key1:value1", "key2:value2"},
			}
			key := datastore.KeyForObj(ctx, b)
			So(datastore.Put(ctx, &BuildInfra{
				ID:    1,
				Build: key,
				Proto: pb.BuildInfra{
					Buildbucket: &pb.BuildInfra_Buildbucket{
						Hostname: "example.com",
					},
				},
			}), ShouldBeNil)
			So(datastore.Put(ctx, &BuildInputProperties{
				ID:    1,
				Build: key,
				Proto: DSStruct{
					Struct: structpb.Struct{
						Fields: map[string]*structpb.Value{
							"input": {
								Kind: &structpb.Value_StringValue{
									StringValue: "input value",
								},
							},
						},
					},
				},
			}), ShouldBeNil)

			Convey("tags", func() {
				p, err := b.ToProto(ctx)
				So(err, ShouldBeNil)
				So(p.Tags, ShouldResembleProto, []*pb.StringPair{
					{
						Key:   "key1",
						Value: "value1",
					},
					{
						Key:   "key2",
						Value: "value2",
					},
				})
				So(b.Proto.Tags, ShouldBeEmpty)
			})

			Convey("infra", func() {
				p, err := b.ToProto(ctx)
				So(err, ShouldBeNil)
				So(p.Infra, ShouldResembleProto, &pb.BuildInfra{
					Buildbucket: &pb.BuildInfra_Buildbucket{
						Hostname: "example.com",
					},
				})
				So(b.Proto.Infra, ShouldBeNil)
			})

			Convey("input properties", func() {
				p, err := b.ToProto(ctx)
				So(err, ShouldBeNil)
				So(p.Input.Properties, ShouldResembleProtoJSON, `{"input": "input value"}`)
				So(b.Proto.Input, ShouldBeNil)
			})

			Convey("output properties", func() {
				So(datastore.Put(ctx, &BuildOutputProperties{
					ID:    1,
					Build: key,
					Proto: DSStruct{
						Struct: structpb.Struct{
							Fields: map[string]*structpb.Value{
								"output": {
									Kind: &structpb.Value_StringValue{
										StringValue: "output value",
									},
								},
							},
						},
					},
				}), ShouldBeNil)
				p, err := b.ToProto(ctx)
				So(err, ShouldBeNil)
				So(p.Output.Properties, ShouldResembleProtoJSON, `{"output": "output value"}`)
				So(b.Proto.Output, ShouldBeNil)
			})
		})
	})
}