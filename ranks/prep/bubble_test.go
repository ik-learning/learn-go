package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countSwaps(a []int32) int {
	count := 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
				i -= 1
				count += 1
				break
			}
		}
	}
	return count
}

func TestCountSwaps(t *testing.T) {
	var tests = []struct {
		input    []int32
		expected int
	}{
		{[]int32{1, 3, 2}, 1},
		{[]int32{4, 2, 3, 1}, 5},
	}

	for _, test := range tests {
		result := countSwaps(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
