package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func alternatingCharacters(s string) int32 {
	result := 0
	previous := rune(-1)
	for _, v := range s {
		if previous == v {
			result += 1
		} else {
			previous = v
		}
	}
	return int32(result)
}

func TestAlternateCharacters(t *testing.T) {
	var tests = []struct {
		input    string
		expected int32
	}{
		{"AAAA", 3},
		{"BBBBB", 4},
		{"ABABABAB", 0},
		{"BABABA", 0},
		{"AAABBB", 4},
	}

	for _, test := range tests {
		result := alternatingCharacters(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
