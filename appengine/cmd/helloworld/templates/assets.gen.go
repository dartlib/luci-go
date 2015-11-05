// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// AUTOGENERATED. DO NOT EDIT.

// Package templates is generated by github.com/luci/luci-go/tools/cmd/assets.
//
// It contains all [*.html] files found in the package as byte arrays.
package templates

// GetAsset returns an asset by its name. Returns nil if no such asset.
func GetAsset(name string) []byte {
	return ([]byte)(files[name])
}

// GetAssetString is version of GetAsset that returns string instead of byte
// slice. Returns empty string if no such asset.
func GetAssetString(name string) string {
	return files[name]
}

var importPath = "github.com/luci/luci-go/appengine/cmd/helloworld/templates"

var files = map[string]string{
	"index.html": "\x3c\x21\x44\x4f\x43\x54\x59\x50\x45\x20\x48\x54\x4d\x4c\x3e\x0a" +
		"\x3c\x68\x74\x6d\x6c\x3e\x0a\x20\x20\x3c\x68\x65\x61\x64\x3e\x0a" +
		"\x20\x20\x20\x20\x3c\x74\x69\x74\x6c\x65\x3e\x45\x78\x61\x6d\x70" +
		"\x6c\x65\x3c\x2f\x74\x69\x74\x6c\x65\x3e\x0a\x20\x20\x3c\x2f\x68" +
		"\x65\x61\x64\x3e\x0a\x20\x20\x3c\x62\x6f\x64\x79\x3e\x0a\x20\x20" +
		"\x20\x20\x7b\x7b\x69\x66\x20\x2e\x48\x61\x73\x55\x73\x65\x72\x7d" +
		"\x7d\x0a\x20\x20\x20\x20\x20\x20\x3c\x70\x3e\x48\x69\x20\x74\x68" +
		"\x65\x72\x65\x2c\x20\x7b\x7b\x2e\x55\x73\x65\x72\x2e\x45\x6d\x61" +
		"\x69\x6c\x7d\x7d\x21\x3c\x2f\x70\x3e\x0a\x20\x20\x20\x20\x20\x20" +
		"\x7b\x7b\x69\x66\x20\x2e\x55\x73\x65\x72\x2e\x50\x69\x63\x74\x75" +
		"\x72\x65\x7d\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x3c\x70\x3e" +
		"\x48\x65\x72\x65\x27\x73\x20\x79\x6f\x75\x72\x20\x70\x69\x63\x74" +
		"\x75\x72\x65\x3a\x20\x3c\x69\x6d\x67\x20\x73\x72\x63\x3d\x22\x7b" +
		"\x7b\x2e\x55\x73\x65\x72\x2e\x50\x69\x63\x74\x75\x72\x65\x7d\x7d" +
		"\x22\x3e\x3c\x2f\x70\x3e\x0a\x20\x20\x20\x20\x20\x20\x7b\x7b\x65" +
		"\x6e\x64\x7d\x7d\x0a\x20\x20\x20\x20\x20\x20\x3c\x61\x20\x68\x72" +
		"\x65\x66\x3d\x22\x7b\x7b\x2e\x4c\x6f\x67\x6f\x75\x74\x55\x52\x4c" +
		"\x7d\x7d\x22\x3e\x4c\x6f\x67\x6f\x75\x74\x3c\x2f\x61\x3e\x0a\x20" +
		"\x20\x20\x20\x7b\x7b\x65\x6c\x73\x65\x7d\x7d\x0a\x20\x20\x20\x20" +
		"\x20\x20\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x7b\x7b\x2e\x4c\x6f" +
		"\x67\x69\x6e\x55\x52\x4c\x7d\x7d\x22\x3e\x4c\x6f\x67\x69\x6e\x3c" +
		"\x2f\x61\x3e\x0a\x20\x20\x20\x20\x7b\x7b\x65\x6e\x64\x7d\x7d\x0a" +
		"\x20\x20\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x3c\x2f\x68\x74\x6d\x6c" +
		"\x3e\x0a",
}
