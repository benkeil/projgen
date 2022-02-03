// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

type Template struct {
	Overrides map[string]string `yaml:"overrides"`
	Steps     []Step            `yaml:"steps"`
}

type Step struct {
	Command  string `yaml:"command"`
	Title    string `yaml:"title"`
	Copy     string `yaml:"copy"`
	CopyAll  bool   `yaml:"copyAll"`
	To       string `yaml:"to"`
	Mkdir    string `yaml:"mkdir"`
	Render   string `yaml:"render"`
	Move     string `yaml:"move"`
	MoveUp   string `yaml:"moveup"`
	Override bool   `yaml:"override"`
}

type Params struct {
	ProjectPath string
	ProjectName string
	DevRoot     string
	VcsProvider string
}
