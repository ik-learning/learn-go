package main
// https://github.com/goccy/go-yaml/issues/607
import (
	"bytes"
	"fmt"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/parser"
)

func main() {
	data := []byte(`# Create a new yaml file with an array of object users
users:
  - name: "John"
    age: 30
    email: "john@example.com"
  - name: "Jane"
    age: 25
    email: "jane@example.com"
  - name: "Bob"
    age: 40
    email: "bob@example.com"
`)

	f, err := parser.ParseBytes(data, parser.ParseComments)
	path, err := yaml.PathString("$.users[*].age")
	if err != nil {
		panic(err)
	}

	value := map[string]any{
		"in_years":  10,
		"in_months": 10 * 12,
	}

	b, err := yaml.Marshal(value)
	if err != nil {
		panic(err)
	}

	err = path.ReplaceWithReader(f, bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	fmt.Println(f.String())
}
