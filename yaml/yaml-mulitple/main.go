package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

// Define a struct matching the YAML structure
type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	SSL  bool   `yaml:"ssl"`
}

func main() {
	// YAML string with multiple documents
	content, err := os.ReadFile("yaml/yaml-mulitple/test.yaml")

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create a YAML decoder for processing multiple documents
	decoder := yaml.NewDecoder(bytes.NewReader(content))

	var configs []Config
	for {
		var config Config
		// Decode one document at a time
		err := decoder.Decode(&config)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Error decoding YAML: %v", err)
		}
		configs = append(configs, config)
	}

	// Print all decoded Config structs
	for i, cfg := range configs {
		fmt.Printf("Config %d: %+v\n", i+1, cfg)
	}
}
