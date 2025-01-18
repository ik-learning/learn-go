package main

import (
	"github.com/goccy/go-yaml/lexer"
)

func main() {
	println("Hello, World!")
	tokens := lexer.Tokenize("foo: bar")
	for _, token := range tokens {
		token.Dump()
	}
}
