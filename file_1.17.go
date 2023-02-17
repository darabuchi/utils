//go:build go1.17 || go1.18 || go1.19 || 1.20
// +build go1.17 go1.18 go1.19 1.20

package utils

import (
	"io/fs"
)

func FileReadWithFs(filename string, fsys fs.FS) (content []byte, err error) {
	buf, err := fs.ReadFile(fsys, filename)
	if err != nil {
		return nil, err
	}
	return buf, err
}
