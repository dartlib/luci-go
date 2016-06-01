// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// AUTOGENERATED. DO NOT EDIT.

// This file is generated by github.com/luci/luci-go/tools/cmd/assets.
//
// It contains tests that ensure that assets embedded into the binary are
// identical to files on disk.

package assets

import (
	"go/build"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestAssets(t *testing.T) {
	t.Parallel()

	pkg, err := build.ImportDir(".", build.FindOnly)
	if err != nil {
		t.Fatalf("can't load package: %s", err)
	}

	fail := false
	for name := range Assets() {
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
		t.Fatalf("run 'go generate' to update assets.gen.go")
	}
}
