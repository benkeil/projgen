// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"
)

type CopyStep struct {
	Params *Params
	Title  string
	Src    string
	Dest   string
}

func (step CopyStep) Execute() error {
	src := filepath.Join(step.Params.ProjectPath, ".projgen", step.Src)
	dest := step.Src
	if step.Dest != "" {
		dest = step.Dest
	}
	fmt.Printf("Src %s to %s\n", step.Src, dest)
	return Copy(src, dest)
}

func Copy(src, dest string) error {
	return CopyArgs(src, dest, "-rv")
}

func CopyT(src, dest string) error {
	return CopyArgs(src, dest, "-rvT")
}

func CopyArgs(src, dest, args string) error {
	_, err := ExecuteCmd(true, "gcp", []string{args, src, dest})
	if err != nil {
		errors.Wrap(err, "could not copy files")
	}
	return nil
}
