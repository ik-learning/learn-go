package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	// "strings"
	"github.com/goccy/go-yaml"
)

func main() {
	content, err := os.ReadFile("yaml/yaml-mulitple-hlm/test.yaml")

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	decoder := yaml.NewDecoder(bytes.NewReader(content))

	var suites []TestSuite
	for {
		var s TestSuite
		if err := decoder.Decode(&s); err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Error decoding YAML: %v", err)
		}
		suites = append(suites, s)
	}

	fmt.Println("line 44", len(suites))
}
