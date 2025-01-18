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
	fmt.Println(string(content))
	decoder := yaml.NewDecoder(bytes.NewReader(content))

	var suites []TestSuite
	for {
		var s TestSuite
		if err := decoder.Decode(&s); err != nil {
			fmt.Println("line 43 error", err)
			break
		}

		suites = append(suites, s)
	}

	fmt.Println("line 44", len(suites))
}
