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
			Src:    step.Copy,
			Dest:   step.To,
		}, nil
	}
	if step.Mkdir != "" {
		return MkdirStep{
			Params: params,
			Path:   step.Mkdir,
		}, nil
	}
	if step.CopyAll == true {
		return CopyAllStep{
			Params: params,
		}, nil
	}
	if step.Render != "" {
		return RenderStep{
			Params: params,
			File:   step.Render,
		}, nil
	}
	if step.Move != "" {
		return MoveStep{
			Params: params,
			From:   step.Move,
			To:     step.To,
		}, nil
	}
	if step.MoveUp != "" {
		return MoveUpStep{
			Params: params,
			From:   step.MoveUp,
		}, nil
	}
	if step.Override {
		return OverrideStep{
			Params: params,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("invalid step configuration: %v", step))
}
