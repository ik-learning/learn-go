package main

import (
	"fmt"
	"log"

	"github.com/vmware-labs/yaml-jsonpath/pkg/yamlpath"
	"gopkg.in/yaml.v3"
)

var data = `
a:
  b:
    - c: 123
  e: |
    Line1
    Line2
`

func main() {
	fmt.Println("legacy yaml")
	var nn yaml.Node
	if err := yaml.Unmarshal([]byte(data), &nn); err != nil {
		log.Fatalf("error: %v", err)
	}
	// works for path "$..a.b[0].c" and "a.b[0].c"
	p, err := yamlpath.NewPath("$..a.b[0].c")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	r, err := p.Find(&nn)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, node := range r {
		fmt.Println(node.Value)
	}
}
