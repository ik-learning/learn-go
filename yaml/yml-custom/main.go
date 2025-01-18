package main
// https://github.com/goccy/go-yaml/issues/475
import (
	"fmt"
	"github.com/goccy/go-yaml"
)

type Sample struct {
	Field string `yaml:"field"`
}

func marshalString(s string) ([]byte, error) {
	return yaml.Marshal(s)
}

func main() {
	sample := Sample{Field: "#foo\a\b\f\v\n\tbar"}
	b, _ := yaml.Marshal(sample)
	fmt.Print(string(b))

	options := []yaml.EncodeOption{yaml.CustomMarshaler[string](marshalString)}
	b, _ = yaml.MarshalWithOptions(sample, options...)
	fmt.Print(string(b))
}
