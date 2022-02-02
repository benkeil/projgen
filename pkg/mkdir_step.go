// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

type MkdirStep struct {
	Params *Params
	Mkdir  string
}

func (step MkdirStep) Execute() error {
	path := filepath.Join(step.Params.ProjectPath, step.Mkdir)
	fmt.Printf("Create directory %s\n", path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
