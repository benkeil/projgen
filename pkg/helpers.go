// Package pkg
// Copyright © 2022 Benedikt Keil <benkeil.me@gmail.com>
package pkg

import (
	"bytes"
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/Masterminds/sprig"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadRemoteTemplateFile(path string) (*Template, error) {
	_, err := os.Stat(path)
	if err != nil {
		r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL: path,
		})
		if err != nil {
			return nil, errors.Wrap(err, "could not clone git repository")
		}
		treeObjects, _ := r.TreeObjects()
		var content string
		_ = treeObjects.ForEach(func(tree *object.Tree) error {
			file, _ := tree.File(".projgen.yaml")
			if file != nil {
				content, _ = file.Contents()
			}
			return nil
		})
		if content == "" {
			return nil, errors.New("could not find .projgen.yaml")
		}
		return readTemplateFile([]byte(content))
	} else {
		return ReadTemplateFile(path)
	}
}

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

func CopyTemplate(url, path string) error {
	err := CopyT(filepath.Join(url), path)
	if err != nil {
		return errors.Wrap(err, "could not copy template")
	}
	_, err = ExecuteCmd(false, "rm", []string{"-rfv", fmt.Sprintf("%s/.git", path)})
	if err != nil {
		return errors.Wrap(err, "could not delete directory")
	}
	return nil
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
	return nil
}

func ReadTemplateFile(path string) (*Template, error) {
	var projgenFile = filepath.Join(path, ".projgen.yaml")
	if _, err := os.Stat(projgenFile); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "directory contains no .projgen.yaml file")
	}
	content, err := ioutil.ReadFile(projgenFile)
	if err != nil {
		return nil, errors.Wrap(err, "could not read .projgen.yaml")
	}
	return readTemplateFile(content)
}

func readTemplateFile(content []byte) (*Template, error) {
	var template Template
	err := yaml.Unmarshal(content, &template)
	if err != nil {
		return nil, errors.Wrap(err, "invalid .projgen.yaml")
	}
	return &template, nil
}

func GitInit(path string) error {
	_, err := ExecuteCmdWorkingDirectory(path, true, "git", "init")
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
	_, err := ExecuteCmdWorkingDirectory(path, true, "git", "add", ".")
	if err != nil {
		return err
	}
	_, err = ExecuteCmdWorkingDirectory(path, true, "git", "commit", "-m", "chore: project created by projgen")
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

func OverrideParam(value string, params Params) string {
	tmpl, err := template.New("override").Funcs(sprig.FuncMap()).Parse(value)
	cobra.CheckErr(err)
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, params)
	return buffer.String()
}
