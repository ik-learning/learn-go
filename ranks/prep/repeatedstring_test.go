package main

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func repeatedString(s string, n int64) int64 {
	remainder := n % int64(len(s))
	multiplier := n / int64(len(s))
	chars := strings.Split(s, "")
	var count int = 0
	for i := 0; i < len(s); i++ {
		if chars[i] == "a" {
			count += 1
		}
	}
	result := int64(count) * multiplier
	for j := 0; j < int(remainder); j++ {
		if chars[j] == "a" {
			result += 1
		}
	}
	return result
}

func TestRepeatedString(t *testing.T) {
	var tests = []struct {
		text       string
		multiplier int64
		expected   int64
	}{
		{"aba", 10, 7},
		{"a", 1000000000000, 1000000000000},
	}

	for _, test := range tests {
		val := repeatedString(test.text, test.multiplier)
		assert.Equal(t, test.expected, val, "they should be equal")
	}
}
