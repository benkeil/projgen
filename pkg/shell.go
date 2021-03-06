// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"fmt"
	"os/exec"
)

func ExecuteCmd(print bool, name string, args []string) (string, error) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	if print {
		fmt.Println(string(stdout))
	}
	return string(stdout), nil
}

func ExecuteCmdWorkingDirectory(workingDirectory string, print bool, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = workingDirectory
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(stdout))
		return "", err
	}
	if print {
		fmt.Println(string(stdout))
	}
	return string(stdout), nil
}
