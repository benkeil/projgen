// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

type Template struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Command string `yaml:"command"`
	Title   string `yaml:"title"`
	Copy    string `yaml:"copy"`
	To      string `yaml:"to"`
	Mkdir   string `yaml:"mkdir"`
}

type Params struct {
	ProjectPath string
	ProjectName string
	DevRoot     string
	VcsProvider string
}
