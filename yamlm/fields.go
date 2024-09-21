package main

import (
	"fmt"
	"log"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Fields map[string]interface{} `yaml:"capabilities"`
}

func allFields(input string) {
	t := Config{}

	err := yaml.Unmarshal([]byte(input), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(t)
	if val, ok := t.Fields["apiVersions"]; ok {
		fmt.Println("\tFound apiVersions")
		if val == nil {
			fmt.Println("\tValue not set")
		} else {
			fmt.Println("\tValue is [", val, "]....", "and type is ", reflect.TypeOf(val))
		}
	}

}
