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

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"go.chromium.org/gae/impl/cloud"
	"go.chromium.org/luci/common/data/caching/lru"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/router"

	"golang.org/x/net/context"
)

// flexBase returns the basic middleware for use on appengine flex.  Flex does not
// allow the use of appengine APIs.
func flexBase() router.MiddlewareChain {
	flex := cloud.Flex{
		Cache: lru.New(1024 * 1024),
	}
	cfg, err := flex.Configure(context.Background())
	if err != nil {
		panic(err)
	}

	// Installs the Info and Datastore services.
	return router.NewMiddlewareChain(func(c *router.Context, next router.Handler) {
		c.Context = cfg.Use(c.Context, flex.Request(c.Request))
		next(c)
	})
}

// Where it all begins!!!
func main() {
	r := router.New()

	base := flexBase()
	r.GET("/log/raw/*path", base, rawLog)

	// Health check, for the appengine flex environment.
	http.HandleFunc("/_ah/health", healthCheckHandler)
	// And everything else.
	http.Handle("/", r)

	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func page(c *router.Context, status int, msg string) {
	c.Writer.WriteHeader(status)
	fmt.Fprintf(c.Writer, msg)
}

func errorPage(c *router.Context, msg string) {
	page(c, http.StatusInternalServerError, msg)
}

func rawLog(c *router.Context) {
	path := c.Params.ByName("path")
	if path == "" {
		page(c, http.StatusBadRequest, "missing path")
		return
	}
	path = strings.TrimLeft(path, "/")
	host := c.Request.FormValue("host")
	if host == "" {
		host = "luci-logdog.appspot.com"
	}
	err := logHandler(c.Context, c.Writer, host, path)
	switch err {
	case nil:
		// Everything is fine
	case errNoAuth:
		// Redirect to login page
		loginURL, err := auth.LoginURL(c.Context, c.Request.URL.Path)
		if err != nil {
			fmt.Fprintf(c.Writer, "Encountered error generating login url: %s\n", err.Error())
			return
		}
		http.Redirect(c.Writer, c.Request, loginURL, http.StatusTemporaryRedirect)
		return
	default:
		fmt.Fprintf(c.Writer, "Encountered error: %s", err.Error())
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
