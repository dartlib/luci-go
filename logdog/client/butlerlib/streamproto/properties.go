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

package streamproto

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"go.chromium.org/luci/common/clock/clockflag"
	"go.chromium.org/luci/common/data/recordio"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/logdog/api/logpb"
	"go.chromium.org/luci/logdog/common/types"
)

// Flags is a flag- and JSON-compatible version of logpb.LogStreamDescriptor.
// It is used for stream negotiation protocol and command-line interfaces.
//
// TODO(iannucci) - Change client->butler protocol to just use jsonpb encoding
// of LogStreamDescriptor.
type Flags struct {
	Name        StreamNameFlag `json:"name,omitempty"`
	ContentType string         `json:"contentType,omitempty"`
	Type        StreamType     `json:"type,omitempty"`
	Timestamp   clockflag.Time `json:"timestamp,omitempty"`
	Tags        TagMap         `json:"tags,omitempty"`
}

// Descriptor converts the Flags to a LogStreamDescriptor.
func (f *Flags) Descriptor() *logpb.LogStreamDescriptor {
	contentType := types.ContentType(f.ContentType)
	if contentType == "" {
		contentType = f.Type.DefaultContentType()
	}

	return &logpb.LogStreamDescriptor{
		Name:        string(f.Name),
		ContentType: string(contentType),
		StreamType:  logpb.StreamType(f.Type),
		Timestamp:   google.NewTimestamp(time.Time(f.Timestamp)),
		Tags:        f.Tags,
	}
}

// WriteHandshake writes the butler protocol header handshake on the given
// Writer.
func (f *Flags) WriteHandshake(w io.Writer) error {
	data, err := json.Marshal(f)
	if err != nil {
		return errors.Annotate(err, "marshaling flags").Err()
	}
	if _, err := w.Write(ProtocolFrameHeaderMagic); err != nil {
		return errors.Annotate(err, "writing magic number").Err()
	}
	if _, err := recordio.WriteFrame(w, data); err != nil {
		return errors.Annotate(err, "writing properties").Err()
	}
	return nil
}

// FromHandshake reads the butler protocol header handshake from the given
// Reader.
func (f *Flags) FromHandshake(r io.Reader) error {
	header := make([]byte, len(ProtocolFrameHeaderMagic))
	_, err := r.Read(header)
	if err != nil {
		return errors.Annotate(err, "reading magic number").Err()
	}
	if !bytes.Equal(header, ProtocolFrameHeaderMagic) {
		return errors.Reason(
			"magic number mismatch: got(%q) expected(%q)",
			header, ProtocolFrameHeaderMagic).Err()
	}
	flagData, err := recordio.NewReader(r, 1024*1024).ReadFrameAll()
	if err != nil {
		return errors.Annotate(err, "reading property frame").Err()
	}
	return errors.Annotate(json.Unmarshal(flagData, f), "parsing flag JSON").Err()
}
