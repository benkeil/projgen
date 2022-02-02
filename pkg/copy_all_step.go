// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"fmt"
	"path/filepath"
)

type CopyAllStep struct {
	Params *Params
}

func (step CopyAllStep) Execute() error {
	src := filepath.Join(step.Params.ProjectPath, ".projgen")
	dest := step.Params.ProjectPath
	fmt.Printf("Src all files to %s\n", dest)
	return CopyT(src, dest)
}
