// Copyright 2017 The LUCI Authors.
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

package upload

import (
	"time"

	"golang.org/x/net/context"

	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"

	api "go.chromium.org/luci/cipd/api/cipd/v1"
)

// Operation is a datastore entity that represents an upload.
type Operation struct {
	_kind  string                `gae:"$kind,cas.UploadOperation"`
	_extra datastore.PropertyMap `gae:"-,extra"`

	ID int64 `gae:"$id"`

	Status api.UploadStatus
	Error  string `gae:",noindex"` // error message if the verification failed

	TempGSPath string `gae:",noindex"` // the GS path to where the client uploads
	UploadURL  string `gae:",noindex"` // resumable upload URL

	HashAlgo  api.HashAlgo // the algo to use to verify the uploaded content
	HexDigest string       // the expected content digest or "" if not known

	CreatedBy identity.Identity // who initiated the upload, FYI
	CreatedTS time.Time         // when the upload was initiated, FYI
	UpdatedTS time.Time         // last time this entity was saved, FYI
}

// ToProto constructs UploadOperation proto message.
//
// The caller must prepare the ID in advance using WrapOpID.
func (op *Operation) ToProto(wrappedID string) *api.UploadOperation {
	var ref *api.ObjectRef
	if op.Status == api.UploadStatus_PUBLISHED {
		ref = &api.ObjectRef{
			HashAlgo:  op.HashAlgo,
			HexDigest: op.HexDigest,
		}
	}
	return &api.UploadOperation{
		OperationId:  wrappedID,
		UploadUrl:    op.UploadURL,
		Status:       op.Status,
		Object:       ref,
		ErrorMessage: op.Error,
	}
}

// Advance transactionally updates the entity (by calling the callback)
// if its Status in the datastore is still same as op.Status.
//
// If the entity in the datastore has different Status, silently skips calling
// the callback. It means the entity has been updated concurrently. This works
// because all Operation mutations actually switch statuses and the Operation
// state machine can "roll" only in one direction.
//
// Returns the most recent state of the operation (whether it was mutated just
// now by the callback or not). 'op' itself is kept intact.
func (op *Operation) Advance(c context.Context, cb func(context.Context, *Operation) error) (*Operation, error) {
	fresh := &Operation{ID: op.ID}
	err := datastore.RunInTransaction(c, func(c context.Context) error {
		if err := datastore.Get(c, fresh); err != nil || fresh.Status != op.Status {
			return err
		}
		if err := cb(c, fresh); err != nil {
			return err
		}
		fresh.UpdatedTS = clock.Now(c).UTC()
		return datastore.Put(c, fresh)
	}, nil)
	if err != nil {
		return nil, errors.Annotate(err, "failed to update the upload operation").Err()
	}
	return fresh, nil
}
