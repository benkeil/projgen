// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"fmt"
	"github.com/pkg/errors"
	"path/filepath"
	"strings"
)

type CopyStep struct {
	Params *Params
	Title  string
	Copy   string
	To     string
}

func (step CopyStep) Execute() error {
	src := filepath.Join(step.Params.ProjectPath, ".projgen", step.Copy)
	dest := step.Copy
	if step.To != "" {
		dest = step.To
	}
	fmt.Printf("Copy %s to %s\n", step.Copy, dest)
	return Copy(src, dest)
}

func Copy(src, dest string) error {
	return CopyArgs(src, dest, "-rv")
}

func CopyA(src, dest string) error {
	return CopyArgs(src, dest, "-rva")
}

func CopyArgs(src, dest, args string) error {
	fmt.Println("gcp", strings.Join([]string{args, src, dest}, " "))
	_, err := ExecuteCmd(true, "gcp", []string{args, src, dest})
	if err != nil {
		errors.Wrap(err, "could not copy files")
	}
	return nil
}
