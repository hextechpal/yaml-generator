package main

import (
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("fixtures/stage_with_10_test.yaml")
	if err != nil {
		log.Fatalf("Error Reading fixture %v", err.Error())
	}

	pipeline := NewPipeline(
		"Parallel150Stages",
		"stages150Parallel",
		"CIProject",
		"default",
		true,
		string(file),
		300)
	err = pipeline.GenerateYaml()
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
}
