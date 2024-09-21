package main

import (
	"fmt"
	"log"

	"sigs.k8s.io/yaml"
)

type Capabilities struct {
	APIVersions *[]string `yaml:"apiVersions"`
}

type Job struct {
	Capabilities Capabilities `yaml:"capabilities"`
}

func k8sYamlCapabilities(input string) {
	t := Job{}

	err := yaml.Unmarshal([]byte(input), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(t)
	fmt.Println(t.Capabilities.APIVersions)
}

func k8sYamlDefault() {
	j := []byte(`{"name": "John", "age": 30}`)
	y, err := yaml.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
	/* Output:
	age: 30
	name: John
	*/
	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(j2))
}
