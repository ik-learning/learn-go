package main

import (
	"fmt"
	"log"
	"strings"

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
	yamlStr := `
host: "example1.com"
port: 443
ssl: true
---
# yaml-language-server: $schema=https://raw.githubusercontent.com/helm-unittest/helm-unittest/refs/heads/main/schema/helm-testsuite.json
host: "example2.com"
port: 80
ssl: false
---
host: "example3.com"
port: 8080
ssl: true
`

	// Create a YAML decoder for processing multiple documents
	decoder := yaml.NewDecoder(strings.NewReader(yamlStr))

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
