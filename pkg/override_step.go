// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"fmt"
	"path/filepath"
)

type OverrideStep struct {
	Params *Params
}

func (step OverrideStep) Execute() error {
	src := filepath.Join(step.Params.ProjectPath, ".projgen")
	dest := filepath.Join(step.Params.ProjectPath, ".")
	fmt.Println("==> Override existing files with project files")
	return CopyArgs(src, dest, "-rvfT")
}
