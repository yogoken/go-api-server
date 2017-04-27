package util

import (
	"go/build"
)

func ImportPkg(path, dir string) *build.Package {

	p, err := build.Import(path, dir, build.ImportComment)
	if err != nil {
		panic(err)
	}

	if p.IsCommand() {
		return nil
	}

	return p
}