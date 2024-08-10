//go:build !unix

// Copyright 2024 Jelly Terra
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package fscopy

func Copy(src, dest string) error { return RawCopy(src, dest) }
