package fizbuzz

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

var tests = []struct {
		input    int
		expected string
	}{
		{12, "Fizz"},
		{4, ""},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}

func TestReverseToReturnReverseInputString(t *testing.T) {
	for _, test := range tests {
		result := fizzBbuzz(test.input)
		assert.Equal(t, result, test.expected, "they should be equal")
	}
}
