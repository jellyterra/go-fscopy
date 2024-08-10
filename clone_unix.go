//go:build unix

// Copyright 2024 Jelly Terra
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package fscopy

import (
	"golang.org/x/sys/unix"
)

func Clone(src, dest string) error {
	var stat unix.Stat_t

	err := unix.Stat(src, &stat)
	if err != nil {
		return err
	}

	srcFd, err := unix.Open(src, unix.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer unix.Close(srcFd)

	dstFd, err := unix.Creat(dest, stat.Mode)
	if err != nil {
		return err
	}
	defer unix.Close(dstFd)

	err = unix.IoctlFileClone(dstFd, srcFd)
	if err != nil {
		return err
	}

	return nil
}
