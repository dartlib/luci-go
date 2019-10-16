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

package internal

import (
	"context"
	"fmt"

	"cloud.google.com/go/compute/metadata"

	"golang.org/x/oauth2/google"

	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/logging"
)

type gceTokenProvider struct {
	account  string
	email    string
	cacheKey CacheKey
}

// NewGCETokenProvider returns TokenProvider that knows how to use GCE metadata server.
func NewGCETokenProvider(ctx context.Context, account string, scopes []string) (TokenProvider, error) {
	// Grab an email associated with the account. This must not be failing on
	// a healthy VM if the account is present. If it does, the metadata server is
	// broken.
	email, err := metadata.Get("instance/service-accounts/" + account + "/email")
	if err != nil {
		if _, yep := err.(metadata.NotDefinedError); yep {
			return nil, ErrInsufficientAccess
		}
		return nil, err
	}

	// Ensure the account has requested scopes. Assume 'cloud-platform' scope
	// covers all possible scopes. This is important when using GKE Workload
	// Identities: the metadata server always reports only 'cloud-platform' scope
	// there. Its presence should be enough to cover all scopes used in practice.
	// The exception is non-cloud scopes (like gerritcodereview or G Suite). To
	// use such scopes, one will have to use impersonation through Cloud IAM APIs,
	// which *are* covered by cloud-platform (see ActAsServiceAccount in auth.go).
	availableScopes, err := metadata.Scopes(account)
	if err != nil {
		return nil, err
	}
	availableSet := stringset.NewFromSlice(availableScopes...)
	if !availableSet.Has("https://www.googleapis.com/auth/cloud-platform") {
		for _, requested := range scopes {
			if !availableSet.Has(requested) {
				logging.Warningf(ctx, "GCE service account %q doesn't have required scope %q (all scopes: %q)", account, requested, availableScopes)
				return nil, ErrInsufficientAccess
			}
		}
	}

	return &gceTokenProvider{
		account: account,
		email:   email,
		cacheKey: CacheKey{
			Key:    fmt.Sprintf("gce/%s", account),
			Scopes: scopes,
		},
	}, nil
}

func (p *gceTokenProvider) RequiresInteraction() bool {
	return false
}

func (p *gceTokenProvider) Lightweight() bool {
	return true
}

func (p *gceTokenProvider) Email() string {
	return p.email
}

func (p *gceTokenProvider) CacheKey(ctx context.Context) (*CacheKey, error) {
	return &p.cacheKey, nil
}

func (p *gceTokenProvider) MintToken(ctx context.Context, base *Token) (*Token, error) {
	src := google.ComputeTokenSource(p.account)
	tok, err := src.Token()
	if err != nil {
		return nil, err
	}
	return &Token{
		Token: *tok,
		Email: p.Email(),
	}, nil
}

func (p *gceTokenProvider) RefreshToken(ctx context.Context, prev, base *Token) (*Token, error) {
	// Minting and refreshing on GCE is the same thing: a call to metadata server.
	return p.MintToken(ctx, base)
}
