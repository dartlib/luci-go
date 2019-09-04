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

package streamserver

import (
	"context"
	"io"

	"go.chromium.org/luci/logdog/api/logpb"
)

// StreamServer is an interface to a background service that allows external
// processes to establish Butler streams.
type StreamServer interface {
	// Performs initial connection and setup, entering a listening state.
	Listen() error

	// Address returns a string that can be used by the "streamclient" package to
	// return a client for this StreamServer.
	//
	// Full package is:
	// go.chromium.org/luci/logdog/butlerlib/streamclient
	Address() string

	// Blocks, returning a new Stream when one is available. If the stream server
	// has closed, this will return nil.
	Next() (io.ReadCloser, *logpb.LogStreamDescriptor)

	// Closes the stream server, cleaning up resources.
	Close()
}

var newStreamServer func(ctx context.Context, path string) (StreamServer, error)

// New creates a new StreamServer.
//
// This has a platform-specific implementation.
//
// On Mac/Linux, this makes a Unix Domain Socket based server, and `path` is
// required to be the absolute path to a filesystem location which is suitable
// for a domain socket. The location will be removed prior to, and after, use.
//
// On Windows, this makes a Named Pipe based server. `path` must be a valid UNC
// path component, and will be used to listen on the named pipe
// "\\.\$path.$PID.$UNIQUE" where $PID is the process id of the current process
// and $UNIQUE is a monotonically increasing value within this process' memory.
//
// `path` may be empty and a unique name will be chosen for you.
func New(ctx context.Context, path string) (StreamServer, error) {
	return newStreamServer(ctx, path)
}
