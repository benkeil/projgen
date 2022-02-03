// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/pkg/errors"
	"html/template"
	"strings"
)

type MoveStep struct {
	Params *Params
	From   string
	To     string
}

func (step MoveStep) Execute() error {
	return Move(step.From, step.To, step.Params)
}

func Move(src, dest string, params *Params) error {
	tmplSrc, err := template.New("src").Funcs(sprig.FuncMap()).Parse(src)
	if err != nil {
		return err
	}
	var bufferSrc bytes.Buffer
	err = tmplSrc.Execute(&bufferSrc, params)
	src = bufferSrc.String()

	tmplDest, err := template.New("dest").Funcs(sprig.FuncMap()).Parse(dest)
	if err != nil {
		return err
	}
	var bufferDest bytes.Buffer
	err = tmplDest.Execute(&bufferDest, params)
	dest = bufferDest.String()

	args := []string{"-av", "--remove-source-files", src, dest}
	fmt.Println(fmt.Sprintf("==> rsync %s", strings.Join(args, " ")))
	_, err = ExecuteCmdWorkingDirectory(params.ProjectPath, true, "rsync", args...)
	if err != nil {
		return errors.Wrap(err, "could not move files")
	}

	args = []string{"-rfv", src}
	fmt.Println(fmt.Sprintf("==> rm %s", strings.Join(args, " ")))
	_, err = ExecuteCmdWorkingDirectory(params.ProjectPath, true, "rm", args...)
	if err != nil {
		return errors.Wrap(err, "could not move files")
	}
	return nil
}
