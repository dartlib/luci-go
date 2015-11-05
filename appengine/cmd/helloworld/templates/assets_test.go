// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// AUTOGENERATED. DO NOT EDIT.

// This file is generated by github.com/luci/luci-go/tools/cmd/assets.
//
// It contains tests that ensure that assets embedded into the binary are
// identical to files on disk.

package templates

import (
	"go/build"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestAssets(t *testing.T) {
	t.Parallel()

	pkg, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		t.Fatalf("can't find package %q", importPath)
	}

	fail := false
	for name := range files {
		GetAsset(name) // for code coverage
		path := filepath.Join(pkg.Dir, filepath.FromSlash(name))
		blob, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("can't read file with assets %q (%s) - %s", name, path, err)
			fail = true
		} else if string(blob) != GetAssetString(name) {
			t.Errorf("embedded asset %q is out of date", name)
			fail = true
		}
	}

	if fail {
		t.Fatalf("run 'go generate %s' to update assets.gen.go", importPath)
	}
}
