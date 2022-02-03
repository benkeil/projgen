// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"fmt"
)

type MoveUpStep struct {
	Params *Params
	From   string
}

func (step MoveUpStep) Execute() error {
	return Move(fmt.Sprintf("%s/", step.From), ".", step.Params)
}
