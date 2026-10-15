// Package utils
package utils

import (
	"errors"
	"path/filepath"
	"runtime"
)

func CurrentDirPath(level ...int) (path string, err error) {
	lvl := 1
	if len(level) > 0 {
		lvl = level[0]
	}
	_, filename, _, ok := runtime.Caller(lvl)
	if !ok {
		return path, errors.New("unable to get caller information")
	}
	path = filepath.Dir(filename)
	return path, nil
}
