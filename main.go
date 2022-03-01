package main

import (
	"log"
)

func main() {
	pipeline := NewPipeline(
		"CIProject",
		"default",
		true,
		200)
	err := pipeline.GenerateYaml()
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
}
