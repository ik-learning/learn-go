package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cache = make(map[int]int)

func fibonacci(val int) int {
	if val == 0 {
		return 0
	} else if val == 1 {
		return 1
	}
	var firstResult int
	var secondResult int

	if first, ok := cache[val - 1]; ok {
		firstResult = first
	} else {
		firstResult = fibonacci(val - 1)
		cache[val - 1] = firstResult
	}
	if second, ok := cache[val - 2]; ok {
		secondResult = second
	} else {
		secondResult = fibonacci(val - 2)
		cache[val - 2] = secondResult
	}

	return firstResult + secondResult
}

func TestFibonacci(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{6, 8},
	}

	for _, test := range tests {
		result := fibonacci(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}

// func BenchmarkCalc(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         _ = fibonacci(i)
//     }
// }
