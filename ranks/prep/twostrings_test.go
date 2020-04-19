package main

import (
	// "strings"

	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.hackerrank.com/challenges/two-strings/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
func twoStrings(s1 string, s2 string) string {
	vals := make(map[byte]bool)
	// for _, c := range []rune(A) {
	// 	amap[c] += 1
	// }
	for i := 0; i < len(s1); i++ {
		vals[s1[i]] = true
	}
	for j := 0; j < len(s2); j++ {
		if _, ok := vals[s2[j]]; ok {
			return "YES"
		}
	}
	return "NO"
}

func TestTwoString(t *testing.T) {
	var tests = []struct {
		first    string
		second   string
		expected string
	}{
		{"ab", "abc", "YES"},
		{"hi", "world", "NO"},
	}

	for _, test := range tests {
		result := twoStrings(test.first, test.second)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
