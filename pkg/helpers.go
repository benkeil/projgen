// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateProject(projectName string) (string, error) {
	_, err := ExecuteCmd(false, "ghq", []string{"create", projectName})
	if err != nil {
		return "", errors.Wrap(err, "could not create git repository")
	}
	path, err := ExecuteCmd(false, "ghq", []string{"list", "-e", "-p", projectName})
	path = strings.Trim(path, "\n")
	if err != nil {
		return "", errors.Wrap(err, "could not get path of created repository")
	}
	_, err = ExecuteCmd(false, "rm", []string{"-rfv", fmt.Sprintf("%s/.git", path)})
	if err != nil {
		return "", errors.Wrap(err, "could not delete directory")
	}
	return path, nil
}

func CloneTemplate(url, path string) error {
	_, err := ExecuteCmd(false, "git", []string{"clone", "--depth=1", url, path})
	if err != nil {
		return errors.Wrap(err, "could not clone template")
	}
	_, err = ExecuteCmd(false, "rm", []string{"-rfv", fmt.Sprintf("%s/.git", path)})
	if err != nil {
		return errors.Wrap(err, "could not delete directory")
	}
	_, err = ExecuteCmd(false, "git", []string{"init", path})
	if err != nil {
		return errors.Wrap(err, "could not init git repository")
	}
	return nil
}

func ReadTemplateFile(path string) (*Template, error) {
	fmt.Printf("create project from template at %v\n", path)
	var projgenFile = filepath.Join(path, ".projgen.yaml")
	if _, err := os.Stat(projgenFile); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "directory contains no .projgen.yaml file")
	}
	yamlFile, err := ioutil.ReadFile(projgenFile)
	if err != nil {
		return nil, errors.Wrap(err, "could not read .projgen.yaml")
	}

	var template Template
	err = yaml.Unmarshal(yamlFile, &template)
	if err != nil {
		return nil, errors.Wrap(err, "invalid .projgen.yaml")
	}
	fmt.Printf("%v\n", template)
	return &template, nil
}

func GitInit(path string) error {
	_, err := ExecuteCmdWorkingDirectory(path, true, "git", []string{"init"})
	if err != nil {
		return err
	}
	return nil
}

func Cleanup(path string) error {
	configFile := filepath.Join(path, ".projgen.yaml")
	configDirectory := filepath.Join(path, ".projgen/")
	fmt.Println("Cleanup", configFile)
	_, err := ExecuteCmd(true, "rm", []string{"-vf", configFile})
	if err != nil {
		return err
	}
	fmt.Println("Cleanup", configDirectory)
	_, err = ExecuteCmd(true, "rm", []string{"-vrf", configDirectory})
	if err != nil {
		return err
	}
	return nil
}

func GitCommit(path string) error {
	_, err := ExecuteCmdWorkingDirectory(path, true, "git", []string{"add", "."})
	if err != nil {
		return err
	}
	_, err = ExecuteCmdWorkingDirectory(path, true, "git", []string{"commit", "-m", "chore: project created by projgen"})
	if err != nil {
		return err
	}
	return nil
}

func Doc(text string) string {
	doc := heredoc.Doc(text)
	var result = ""
	for _, line := range strings.Split(doc, "\n") {
		result += fmt.Sprint("  ", line, "\n")
	}
	return strings.Trim(result, "\n")
}
