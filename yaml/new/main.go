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

var doc = `
s:
  - a:
      b: u1
      c: get1
      d: i1
  - w:
      c: bad
      e:
        - a:
            b: u2
            c: get2
            d: i2
`

func main() {
	fmt.Println("New Library")
	path, err := yaml.PathString("$.s[*].a.c")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("path:", path.String())
	n, err := path.ReadNode(strings.NewReader(doc))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("value:\n", n)
}
