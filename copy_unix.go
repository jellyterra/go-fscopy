//go:build unix

// Copyright 2024 Jelly Terra
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package fscopy

import (
	"errors"
	"syscall"
)

func Copy(src, dest string) error {
	err := Clone(src, dest)
	switch {
	case err == nil:
		return nil
	case errors.Is(err.(syscall.Errno), syscall.EOPNOTSUPP):
		fallthrough
	case errors.Is(err.(syscall.Errno), syscall.EXDEV):
		return RawCopy(src, dest)
	default:
		return err
	}
}
