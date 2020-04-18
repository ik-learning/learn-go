package main

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

// https://www.hackerrank.com/challenges/sparse-arrays/problem
func matchingStrings(strings []string, queries []string) []int32 {
	counter := make(map[string]int32)
	for _, val := range strings {
		if _, ok := counter[val]; ok {
			counter[val]++
		} else {
			counter[val] = 1
		}
	}
	result := make([]int32, len(queries))
	for i, val := range queries {
		if el, ok := counter[val]; ok {
			result[i] = el
		}
	}
	return result
}

func TestMatchStrings(t *testing.T) {
	var tests = []struct {
		input    []string
		query    []string
		expected []int32
	}{
		{[]string{"def", "de", "fgh"}, []string{"de", "lmn", "fgh"}, []int32{1, 0, 1}},
		{[]string{
			"aba",
			"baba",
			"aba",
			"xzxb",
		},
			[]string{"aba", "xzxb", "ab"}, []int32{2, 1, 0}},
		{[]string{
			"abcde",
			"sdaklfj",
			"asdjf",
			"na",
			"basdn",
			"sdaklfj",
			"asdjf",
			"na",
			"asdjf",
			"na",
			"basdn",
			"sdaklfj",
			"asdjf",
		},
			[]string{
				"abcde",
				"sdaklfj",
				"asdjf",
				"na",
				"basdn",
			}, []int32{1, 3, 4, 3, 2}},
	}

	for _, test := range tests {
		result := matchingStrings(test.input, test.query)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
