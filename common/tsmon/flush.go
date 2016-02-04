// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package tsmon

import (
	"errors"
	"time"

	"golang.org/x/net/context"

	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/logging"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Flush sends all the metrics that are registered in the application.
func Flush(ctx context.Context) error {
	mon := Monitor()
	tar := Target()
	if mon == nil || tar == nil {
		return errors.New("no tsmon Monitor is configured")
	}

	// Split up the payload into chunks if there are too many cells.
	cells := Store().GetAll(ctx)

	chunkSize := mon.ChunkSize()
	if chunkSize == 0 {
		chunkSize = len(cells)
	}
	for len(cells) > 0 {
		count := minInt(chunkSize, len(cells))
		if err := mon.Send(cells[:count], tar); err != nil {
			return err
		}
		cells = cells[count:]
	}
	return nil
}

// autoFlusher knows how to periodically call 'Flush'.
type autoFlusher struct {
	killed chan struct{}
	cancel context.CancelFunc

	flush func(context.Context) error // mocked in unit tests
}

func (f *autoFlusher) start(ctx context.Context, interval time.Duration) {
	flush := f.flush
	if flush == nil {
		flush = Flush
	}

	// 'killed' is closed when timer goroutine exits.
	killed := make(chan struct{})
	f.killed = killed

	ctx, f.cancel = context.WithCancel(ctx)
	go func() {
		defer close(killed)

		for {
			select {
			case <-ctx.Done():
				return
			case <-clock.After(ctx, interval):
				if err := flush(ctx); err != nil && err != context.Canceled {
					logging.Warningf(ctx, "Failed to flush tsmon metrics: %v", err)
				}
			}
		}
	}()
}

func (f *autoFlusher) stop() {
	f.cancel()
	<-f.killed
	f.cancel = nil
	f.killed = nil
}
