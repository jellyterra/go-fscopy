// Copyright 2024 Jelly Terra
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package fscopy

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func RawCopy(src, dest string) error {

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	srcStat, err := srcFile.Stat()
	if err != nil {
		return err
	}

	dstFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, srcStat.Mode())
	if err != nil {
		return err
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func CopyAll(src, dest string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, _ := filepath.Rel(src, path)
		destPath := filepath.Join(dest, relativePath)

		if info.IsDir() {
			return os.Mkdir(destPath, info.Mode())
		}

		err = Copy(path, destPath)
		if err != nil {
			return err
		}

		return nil
	})
}

func CopyAllWithExceptionGlobs(src, dest string, except ...*regexp.Regexp) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, _ := filepath.Rel(src, path)
		destPath := filepath.Join(dest, relativePath)

		for _, glob := range except {
			if glob.MatchString(relativePath) {
				return nil
			}
		}

		if info.IsDir() {
			return os.Mkdir(destPath, info.Mode())
		}

		err = Copy(path, destPath)
		if err != nil {
			return err
		}

		return nil
	})
}
