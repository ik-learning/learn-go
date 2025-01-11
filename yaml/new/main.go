package main

import (
	"fmt"
	"log"
	"strings"

	yaml "github.com/goccy/go-yaml"
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
	path, err := yaml.PathString("$..a.b[0].c")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("path:", path.String())
	n, err := path.ReadNode(strings.NewReader(data))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("value:\n", n)
}
