package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type Pipeline struct {
	Name       string
	Identifier string
	Project    string
	Org        string
	IsParallel bool
	Sl         []int
}

func NewPipeline(project string, org string, isParallel bool, stageCount int) *Pipeline {
	sl := make([]int, stageCount)
	for i := 0; i < stageCount; i++ {
		sl[i] = i
	}
	suffix := ""
	if isParallel {
		suffix = "parallel"
	}
	name := fmt.Sprintf("Stages%d%s", stageCount, suffix)
	return &Pipeline{
		Name:       name,
		Identifier: name,
		Project:    project,
		Org:        org,
		IsParallel: isParallel,
		Sl:         sl,
	}
}

func (p *Pipeline) GenerateYaml() error {
	in, err := ioutil.ReadFile("pipeline.yaml")
	if err != nil {
		return err
	}
	t := template.Must(template.New("tmpl").Parse(string(in)))
	f, err := os.Create(fmt.Sprintf("output/%s.yaml", p.Name))
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, p)
}
