// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package model

import (
	"testing"

	"golang.org/x/net/context"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/luci/gae/filter/featureBreaker"
	"github.com/luci/gae/impl/memory"
	"github.com/luci/gae/service/datastore"

	"github.com/luci/luci-go/common/clock/testclock"
	"github.com/luci/luci-go/common/logging/memlogger"
	"github.com/luci/luci-go/common/proto/google"
	. "github.com/luci/luci-go/common/testing/assertions"

	"github.com/luci/luci-go/common/api/dm/service/v1"
)

func TestExecutions(t *testing.T) {
	t.Parallel()

	Convey("Execution", t, func() {
		c := memory.Use(context.Background())
		c = memlogger.Use(c)
		ds := datastore.Get(c)

		a := &Attempt{ID: *dm.NewAttemptID("q", 1)}
		ak := ds.KeyForObj(a)

		Convey("Revoke", func() {
			e1 := &Execution{ID: 1, Attempt: ak, Token: []byte("good tok")}
			So(ds.Put(e1), ShouldBeNil)

			e2 := *e1
			So(e2.Revoke(c), ShouldBeNil)

			So(e1.Token, ShouldResemble, []byte("good tok"))
			So(ds.Get(e1), ShouldBeNil)
			So(e1.Token, ShouldBeNil)
		})

		Convey("Verify", func() {
			e1 := &Execution{ID: 1, Attempt: ak, Token: []byte("good tok")}
			So(ds.Put(e1), ShouldBeNil)

			auth := &dm.Execution_Auth{
				Id:    dm.NewExecutionID("q", a.ID.Id, uint32(e1.ID)),
				Token: []byte("bad tok"),
			}

			_, _, err := AuthenticateExecution(c, auth)
			So(err, ShouldBeRPCUnauthenticated, "execution Auth")

			So(ds.Put(a), ShouldBeNil)
			_, _, err = AuthenticateExecution(c, auth)
			So(err, ShouldBeRPCUnauthenticated, "execution Auth")

			a.CurExecution = 1
			So(ds.Put(a), ShouldBeNil)
			_, _, err = AuthenticateExecution(c, auth)
			So(err, ShouldBeRPCUnauthenticated, "execution Auth")

			a.State = dm.Attempt_EXECUTING
			So(ds.Put(a), ShouldBeNil)
			_, _, err = AuthenticateExecution(c, auth)
			So(err, ShouldBeRPCUnauthenticated, "execution Auth")

			e1.State = dm.Execution_RUNNING
			So(ds.Put(e1), ShouldBeNil)
			_, _, err = AuthenticateExecution(c, auth)
			So(err, ShouldBeRPCUnauthenticated, "execution Auth")

			auth.Token = []byte("good tok")
			atmpt, exe, err := AuthenticateExecution(c, auth)
			So(err, ShouldBeNil)

			So(atmpt, ShouldResemble, a)
			So(exe, ShouldResemble, e1)
		})

		Convey("Activate", func() {
			e1 := &Execution{
				ID:      1,
				Attempt: ak,
				Token:   []byte("good tok"),
			}
			a.CurExecution = 1
			So(ds.PutMulti([]interface{}{a, e1}), ShouldBeNil)

			auth := &dm.Execution_Auth{
				Id:    dm.NewExecutionID("q", a.ID.Id, uint32(e1.ID)),
				Token: []byte("wrong tok"),
			}

			Convey("wrong execution id", func() {
				auth.Id.Id++
				_, _, err := ActivateExecution(c, auth, []byte("wrong tok"))
				So(err, ShouldBeRPCUnauthenticated, "execution Auth")
			})

			Convey("attempt bad state", func() {
				_, _, err := ActivateExecution(c, auth, []byte("wrong tok"))
				So(err, ShouldBeRPCUnauthenticated, "execution Auth")
			})

			Convey("attempt executing", func() {
				a.State = dm.Attempt_EXECUTING
				So(ds.Put(a), ShouldBeNil)

				Convey("wrong execution state", func() {
					e1.State = dm.Execution_CANCELLED
					So(ds.Put(e1), ShouldBeNil)
					_, _, err := ActivateExecution(c, auth, []byte("wrong token"))
					So(err, ShouldBeRPCUnauthenticated, "execution Auth")
				})

				Convey("wrong token", func() {
					_, _, err := ActivateExecution(c, auth, []byte("wrong tok"))
					So(err, ShouldBeRPCUnauthenticated, "execution Auth")
				})

				Convey("correct token", func() {
					auth.Token = []byte("good tok")
					memlogger.Reset(c)
					newA, e, err := ActivateExecution(c, auth, []byte("new token"))
					memlogger.MustDumpStdout(c)
					So(err, ShouldBeNil)
					So(newA, ShouldResemble, a)
					So(e.State, ShouldEqual, dm.Execution_RUNNING)

					Convey("retry with different token fails", func() {
						_, _, err = ActivateExecution(c, auth, []byte("other token"))
						So(err, ShouldBeRPCUnauthenticated, "execution Auth")
					})

					Convey("retry with same token OK", func() {
						auth.Token = []byte("new token")
						_, _, err = ActivateExecution(c, auth, []byte("new token"))
						So(err, ShouldBeNil)
					})
				})
			})

		})

		Convey("Invalidate", func() {
			e1 := &Execution{
				ID:      1,
				Attempt: ak,
				Token:   []byte("good tok"),
				State:   dm.Execution_RUNNING,
			}
			So(ds.Put(e1), ShouldBeNil)

			a.CurExecution = 1
			a.State = dm.Attempt_EXECUTING
			So(ds.Put(a), ShouldBeNil)

			auth := &dm.Execution_Auth{
				Id:    dm.NewExecutionID("q", a.ID.Id, uint32(e1.ID)),
				Token: []byte("bad token"),
			}

			_, _, err := InvalidateExecution(c, auth)
			So(err, ShouldBeRPCUnauthenticated, "execution Auth")

			auth.Token = []byte("good tok")
			_, _, err = InvalidateExecution(c, auth)
			So(err, ShouldBeNil)

			So(ds.Get(e1), ShouldBeNil)
			So(e1.Token, ShouldBeNil)

			_, _, err = InvalidateExecution(c, auth)
			So(err, ShouldBeRPCUnauthenticated, "requires execution Auth")
		})

		Convey("failed invalidation", func() {
			e1 := &Execution{
				ID:      1,
				Attempt: ak,
				Token:   []byte("good tok"),
				State:   dm.Execution_RUNNING,
			}
			So(ds.Put(e1), ShouldBeNil)
			a.CurExecution = 1
			a.State = dm.Attempt_EXECUTING
			So(ds.Put(a), ShouldBeNil)

			auth := &dm.Execution_Auth{
				Id:    dm.NewExecutionID("q", a.ID.Id, uint32(e1.ID)),
				Token: []byte("good tok"),
			}

			c, fb := featureBreaker.FilterRDS(c, nil)
			fb.BreakFeatures(nil, "PutMulti")

			_, _, err := InvalidateExecution(c, auth)
			So(err, ShouldBeRPCInternal, "unable to invalidate Auth")

			fb.UnbreakFeatures("PutMulti")

			_, _, err = InvalidateExecution(c, auth)
			So(err, ShouldBeNil)

			So(ds.Get(e1), ShouldBeNil)
			So(e1.Token, ShouldBeNil)
		})
	})

}

func TestExecutionToProto(t *testing.T) {
	t.Parallel()

	Convey("Test Execution.ToProto", t, func() {
		c := memory.Use(context.Background())
		c = memlogger.Use(c)
		ds := datastore.Get(c)

		e := &Execution{
			ID:      1,
			Attempt: ds.MakeKey("Attempt", "qst|fffffffe"),

			StateReason: "scheduled by DM",

			Created:          testclock.TestTimeUTC,
			DistributorToken: "id",
			DistributorURL:   "https://thing.place.example.com/task/id",

			Token: []byte("secret"),
		}

		Convey("no id", func() {
			So(e.ToProto(false), ShouldResemble, &dm.Execution{Data: &dm.Execution_Data{
				State:              dm.Execution_SCHEDULED,
				StateReason:        "scheduled by DM",
				Created:            google.NewTimestamp(testclock.TestTimeUTC),
				DistributorToken:   "id",
				DistributorInfoUrl: "https://thing.place.example.com/task/id",
			}})
		})

		Convey("with id", func() {
			So(e.ToProto(true), ShouldResemble, &dm.Execution{
				Id: dm.NewExecutionID("qst", 1, 1),
				Data: &dm.Execution_Data{
					State:              dm.Execution_SCHEDULED,
					StateReason:        "scheduled by DM",
					Created:            google.NewTimestamp(testclock.TestTimeUTC),
					DistributorToken:   "id",
					DistributorInfoUrl: "https://thing.place.example.com/task/id",
				},
			})
		})
	})
}
