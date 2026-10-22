package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

func GetGoModulePath() string {
	path, err := ReadGoModulePath(".")
	if err != nil {
		panic(err)
	}

	return path
}

// ReadGoModulePath reads the module path from the go.mod file in dir.
func ReadGoModulePath(dir string) (string, error) {
	path := filepath.Join(dir, "go.mod")
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}

	mod, err := modfile.Parse(path, data, nil)
	if err != nil {
		return "", fmt.Errorf("parse %s: %w", path, err)
	}
	if mod.Module == nil {
		return "", errors.New("go.mod does not declare a module path")
	}

	return mod.Module.Mod.Path, nil
}
