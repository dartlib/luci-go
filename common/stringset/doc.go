// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Package stringset is an exceedingly simple 'set' implementation for strings.
//
// It's not threadsafe, but can be used in place of a simple
// `map[string]struct{}`
package stringset
