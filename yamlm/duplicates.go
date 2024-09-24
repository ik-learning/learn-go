package main

// https://github.com/go-yaml/yaml/issues/369
import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type B struct{}
type A struct {
	ZeName string `yaml:"boo"`
	Boo    *B     `yaml:"-"`
}

func duplicates() {
	a := A{}

	_ = yaml.Unmarshal([]byte("{boo: bar}"), &a)
	fmt.Println(a.Boo)

}

func (t *A) UnmarshalYAML(node *yaml.Node) error {
	fmt.Println("------> unmarshal yaml")
	fmt.Println("------> unmarshal yaml")
	type alias A
	dataAlias := (*alias)(t)

	err := node.Decode(&dataAlias)
	if err != nil {
		return err
	}
	// d.AA = "AA Test"
	return nil
}
