package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Fields map[string]string `yaml:"capabilities"`
}

func allFields(input string) {
	t := Config{}

	err := yaml.Unmarshal([]byte(input), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(t)
}
