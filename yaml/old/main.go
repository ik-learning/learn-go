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

var yml2 = `
store:
  - book:
    - author: john
      price: 10
    - author: ken
      price: 12
  - bicycle:
    color: red
    price: 19.95
  - bicycle*unicycle:
    price: 20.25
`

func main() {
	fmt.Println("legacy yaml")
	var nn yaml.Node
	if err := yaml.Unmarshal([]byte(doc), &nn); err != nil {
		log.Fatalf("error: %v", err)
	}
	// works for path "$..a.b[0].c" and "a.b[0].c"
	p, err := yamlpath.NewPath("$..w.e[*].a")
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
