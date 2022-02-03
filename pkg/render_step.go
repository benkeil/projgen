package pkg

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"html/template"
	"os"
	"path/filepath"
)

type RenderStep struct {
	Params *Params
	File   string
}

func (step RenderStep) Execute() error {
	file := filepath.Join(step.Params.ProjectPath, step.File)
	fmt.Println(fmt.Sprintf("==> Render %s %+v", step.File, step.Params))
	stat, err := os.Stat(file)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	tmpl, err := template.New("render").Funcs(sprig.FuncMap()).Parse(string(data))
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, step.Params)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, buffer.Bytes(), stat.Mode())
	if err != nil {
		return err
	}
	return nil
}
