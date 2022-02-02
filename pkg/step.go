// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"errors"
	"fmt"
)

type ExecutableStep interface {
	Execute() error
}

func (step Step) Transform(params *Params) (ExecutableStep, error) {
	if step.Command != "" {
		return CommandStep{
			Params:  params,
			Command: step.Command,
			Title:   step.Title,
		}, nil
	}
	if step.Copy != "" {
		return CopyStep{
			Params: params,
			Title:  step.Title,
			Copy:   step.Copy,
			To:     step.To,
		}, nil
	}
	if step.Mkdir != "" {
		return MkdirStep{
			Params: params,
			Mkdir:  step.Mkdir,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid step configuration: %v", step))
}
