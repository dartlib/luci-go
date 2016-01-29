// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package archive constructs a LogDog archive out of log stream components.
// Records are read from the stream and emitted as an archive.
package archive

import (
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/common/parallel"
	"github.com/luci/luci-go/common/proto/logdog/logpb"
	"github.com/luci/luci-go/common/recordio"
)

// Manifest is a set of archival parameters.
type Manifest struct {
	// Desc is the logpb.LogStreamDescriptor for the stream.
	Desc *logpb.LogStreamDescriptor
	// Source is the LogEntry source for the stream.
	Source LogEntrySource

	// LogWriter, if not nil, is the Writer to which the log stream record stream
	// will be written.
	LogWriter io.Writer
	// IndexWriter, if not nil, is the Writer to which the log stream Index
	// protobuf stream will be written.
	IndexWriter io.Writer
	// DataWriter, if not nil, is the Writer to which reconstructed LogEntry data
	// will be written.
	DataWriter io.Writer

	// StreamIndexRange, if >0, is the maximum number of log entry stream indices
	// in between successive index entries.
	//
	// If no index constraints are set, an index entry will be emitted for each
	// LogEntry.
	StreamIndexRange int
	// PrefixIndexRange, if >0, is the maximum number of log entry prefix indices
	// in between successive index entries.
	PrefixIndexRange int
	// ByteRange, if >0, is the maximum number of log entry bytes in between
	// successive index entries.
	ByteRange int

	// Logger, if not nil, will be used to log status during archival.
	Logger logging.Logger

	// sizeFunc is a size method override used for testing.
	sizeFunc func(proto.Message) int
}

func (m *Manifest) logger() logging.Logger {
	if l := m.Logger; l != nil {
		return l
	}
	return logging.Null()
}

// Archive performs the log archival described in the supplied Manifest.
func Archive(m Manifest) error {
	// Wrap our log source in a safeLogEntrySource to protect our index order.
	m.Source = &safeLogEntrySource{
		Manifest:       &m,
		LogEntrySource: m.Source,
	}

	// If no constraints are applied, index every LogEntry.
	if m.StreamIndexRange <= 0 && m.PrefixIndexRange <= 0 && m.ByteRange <= 0 {
		m.StreamIndexRange = 1
	}

	// If we're constructing an index, allocate a stateful index builder.
	var idx *indexBuilder
	if m.IndexWriter != nil {
		idx = &indexBuilder{
			Manifest: &m,
			index: logpb.LogIndex{
				Desc: m.Desc,
			},
			sizeFunc: m.sizeFunc,
		}
	}

	return parallel.FanOutIn(func(taskC chan<- func() error) {
		var logC chan *logpb.LogEntry
		if m.LogWriter != nil {
			logC = make(chan *logpb.LogEntry)
			defer close(logC)

			taskC <- func() error {
				if err := archiveLogs(m.LogWriter, m.Desc, logC, idx); err != nil {
					return err
				}

				// If we're building an index, emit it now that the log stream has
				// finished.
				if idx != nil {
					return idx.emit(m.IndexWriter)
				}
				return nil
			}
		}

		var dataC chan *logpb.LogEntry
		if m.DataWriter != nil {
			dataC = make(chan *logpb.LogEntry)
			defer close(dataC)

			taskC <- func() error {
				return archiveData(m.DataWriter, dataC)
			}
		}

		// Iterate through all of our Source's logs and process them.
		for le := m.Source.NextLogEntry(); le != nil; le = m.Source.NextLogEntry() {
			if logC != nil {
				logC <- le
			}
			if dataC != nil {
				dataC <- le
			}
		}
	})
}

func archiveLogs(w io.Writer, d *logpb.LogStreamDescriptor, logC <-chan *logpb.LogEntry, idx *indexBuilder) error {
	offset := int64(0)
	out := func(pb proto.Message) error {
		d, err := proto.Marshal(pb)
		if err != nil {
			return err
		}

		count, err := recordio.WriteFrame(w, d)
		offset += int64(count)
		return err
	}

	// Start with our descriptor protobuf. Defer error handling until later, as
	// we are still responsible for draining "logC".
	err := out(d)
	for le := range logC {
		if err != nil {
			continue
		}

		// Add this LogEntry to our index, noting the current offset.
		if idx != nil {
			idx.addLogEntry(le, offset)
		}
		err = out(le)
	}
	return err
}
