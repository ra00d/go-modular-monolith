package utils

import (
	"os"

	"golang.org/x/mod/modfile"
)

func GetGoModulePath() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}

	mod, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		panic(err)
	}

	return mod.Module.Mod.Path
}
