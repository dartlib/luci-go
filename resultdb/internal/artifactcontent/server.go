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

package artifactcontent

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/grpc/appstatus"
	"go.chromium.org/luci/grpc/grpcutil"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/server/tokens"

	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/pbutil"
)

var artifactNameTokenKind = tokens.TokenKind{
	Algo:       tokens.TokenAlgoHmacSHA256,
	Expiration: time.Hour,
	SecretKey:  "artifact_name",
	Version:    1,
}

// Server can serve artifact content, and generate signed URLs to the content.
type Server struct {
	// Use http:// (not https://) for generated URLs.
	InsecureURLs bool

	// Included in generated signed URLs and required in content requests.
	Hostname string

	// used for isolate client
	anonClient, authClient *http.Client

	// mock for isolate fetching
	testFetchIsolate func(ctx context.Context, isolateURL string, w io.Writer) error
}

// NewServer creates a Server.
func NewServer(ctx context.Context, insecureURLs bool, hostname string) (*Server, error) {
	anonTransport, err := auth.GetRPCTransport(ctx, auth.NoAuth)
	if err != nil {
		return nil, err
	}
	selfTransport, err := auth.GetRPCTransport(ctx, auth.AsSelf)
	if err != nil {
		return nil, err
	}

	return &Server{
		InsecureURLs: insecureURLs,
		Hostname:     hostname,
		anonClient:   &http.Client{Transport: anonTransport},
		authClient:   &http.Client{Transport: selfTransport},
	}, nil
}

// InstallHandlers installs handlers to serve artifact content.
func (s *Server) InstallHandlers(r *router.Router) {
	// TODO(nodir): use OAuth2.0 middleware to allow OAuth credentials.

	// Ideally we use a more narrow pattern, but we cannot because of
	// https://github.com/julienschmidt/httprouter/issues/208
	// This is triggered by URL-escaped test IDs.
	r.GET("/invocations/*rest", router.NewMiddlewareChain(), s.handleContent)
}

func (s *Server) handleContent(c *router.Context) {
	req := &contentRequest{Server: s, w: c.Writer}
	req.handle(c)
}

type contentRequest struct {
	*Server
	w http.ResponseWriter

	artifactName string

	invID      span.InvocationID
	parentID   string
	artifactID string

	contentType spanner.NullString
	size        spanner.NullInt64
}

func (r *contentRequest) handle(c *router.Context) {
	if err := r.parseRequest(c.Context, c.Request); err != nil {
		r.sendError(c.Context, appstatus.BadRequest(err))
		return
	}

	if err := r.checkAccess(c.Context, c.Request); err != nil {
		r.sendError(c.Context, err)
		return
	}

	// Read the state from database.
	var isolateURL spanner.NullString
	txn := span.Client(c.Context).Single()
	key := r.invID.Key(r.parentID, r.artifactID)
	err := span.ReadRow(c.Context, txn, "Artifacts", key, map[string]interface{}{
		"ContentType": &r.contentType,
		"Size":        &r.size,
		"IsolateURL":  &isolateURL,
	})

	// Check the error and write content to the response body.
	switch {
	case spanner.ErrCode(err) == codes.NotFound:
		err = appstatus.Attachf(err, codes.NotFound, "%s not found", r.artifactName)
		r.sendError(c.Context, err)

	case err != nil:
		r.sendError(c.Context, err)

	case isolateURL.Valid:
		r.handleIsolateContent(c.Context, isolateURL.StringVal)

	default:
		r.sendError(c.Context, appstatus.Errorf(codes.Unimplemented, "currently only isolated artifacts can be served"))
	}
}

func (r *contentRequest) parseRequest(ctx context.Context, req *http.Request) error {
	// We should not use URL.Path because it is important to preserve escaping
	// of test IDs.
	r.artifactName = strings.Trim(req.URL.EscapedPath(), "/")

	invID, testID, resultID, artifactID, err := pbutil.ParseArtifactName(r.artifactName)
	if err != nil {
		return errors.Annotate(err, "invalid artifact name %q", r.artifactName).Err()
	}
	r.invID = span.InvocationID(invID)
	r.parentID = span.ArtifactParentID(testID, resultID)
	r.artifactID = artifactID
	return nil
}

// checkAccess ensures that the requester has access to the artifact content.
//
// If the URL is signed, checks access using token query string param.
// Otherwise, uses OAuth 2.0.
func (r *contentRequest) checkAccess(ctx context.Context, req *http.Request) error {
	token := req.URL.Query().Get("token")
	if token == "" {
		// TODO(nodir): fallback to OAuth 2.0.
		return appstatus.Errorf(codes.Unauthenticated, "no token")
	}

	_, err := artifactNameTokenKind.Validate(ctx, token, []byte(r.artifactName))
	if !transient.Tag.In(err) {
		return appstatus.Attachf(err, codes.PermissionDenied, "invalid token")
	}
	return err
}

func (r *contentRequest) sendError(ctx context.Context, err error) {
	if err == nil {
		panic("err is nil")
	}
	st, ok := appstatus.Get(err)
	httpCode := grpcutil.CodeStatus(st.Code())
	if !ok || httpCode == http.StatusInternalServerError {
		logging.Errorf(ctx, "responding with: %s", err)
		http.Error(r.w, "Internal server error", http.StatusInternalServerError)
	} else {
		logging.Warningf(ctx, "responding with: %s", st.Message())
		http.Error(r.w, st.Message(), httpCode)
	}
}

func (r *contentRequest) writeContentHeaders() {
	if r.contentType.Valid {
		r.w.Header().Set("Content-Type", r.contentType.StringVal)
	}
	if r.size.Valid {
		r.w.Header().Set("Content-Length", strconv.FormatInt(r.size.Int64, 10))
	}
}

// GenerateSignedURL generates a signed HTTPS URL back to this server.
// The returned token works only with the same artifact name.
func (s *Server) GenerateSignedURL(ctx context.Context, artifactName string) (url string, expiration time.Time, err error) {
	now := clock.Now(ctx).UTC()

	tok, err := artifactNameTokenKind.Generate(ctx, []byte(artifactName), nil, artifactNameTokenKind.Expiration)
	if err != nil {
		return "", time.Time{}, err
	}

	scheme := "https"
	if s.InsecureURLs {
		scheme = "http"
	}
	// Using url.URL here is hard because it escapes artifact name which we don't want.
	url = fmt.Sprintf("%s://%s/%s?token=%s", scheme, s.Hostname, artifactName, tok)
	expiration = now.Add(artifactNameTokenKind.Expiration)
	return
}