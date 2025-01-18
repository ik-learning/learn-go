package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	// "strings"
	gyaml "github.com/goccy/go-yaml"
)

func main() {
	content, err := os.ReadFile("yaml/yaml-mulitple-hlm/test.yaml")

	options := []gyaml.DecodeOption{gyaml.CustomUnmarshaler(func(dst *interface{}, b []byte) error {
		if err := gyaml.Unmarshal(b, dst); err != nil {
			return err
		}

		fmt.Println("BINGO", *dst, reflect.TypeOf(*dst))
		// required for backward compoatiblity with gopkg.in/yaml.v3 as it unmarshals all numbers as int64
		if reflect.ValueOf(*dst).CanUint() {
			// fmt.Println("BINGO before", reflect.TypeOf(*dst))
			*dst = int(reflect.ValueOf(*dst).Uint())
			fmt.Println("BINGO after", *dst, reflect.TypeOf(*dst))
			// fmt.Println("BINGO after", *dst, reflect.TypeOf(*dst))
		}
		return nil
	})}

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	decoder := yaml.NewDecoder(bytes.NewReader(content), options...)

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
