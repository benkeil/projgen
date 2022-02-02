// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
)

type CommandStep struct {
	Params  *Params
	Title   string
	Command string
}

func (step CommandStep) Execute() error {
	var title = step.Title
	tmpl, err := template.New("command").Parse(step.Command)
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, step.Params)
	if err != nil {
		return err
	}
	command := buffer.String()
	if title == "" {
		title = command
	}
	fmt.Printf("Execute '%s'\n", title)
	parts := regexp.MustCompile("\\s+").Split(command, -1)
	_, err = ExecuteCmdWorkingDirectory(step.Params.ProjectPath, true, parts[0], parts[1:])
	if err != nil {
		return err
	}
	return nil
}
