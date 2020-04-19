package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxSubsetSum(arr []int32) int32 {
	result := int32(-100000)
	for i, val := range arr {
		tmp := int32(0)
		for j := i + 2; j < len(arr); {
			tmp += arr[j]
			j += 2
		}

		if i+2 < len(arr) && tmp+val > result {
			result = tmp + val
		}

		for j := i + 2; j < len(arr); {
			if result < val+arr[j] {
				result = val + arr[j]
			}
			j += 1
		}
	}
	return int32(result)
}

func TestMaxSubsetSum(t *testing.T) {
	var tests = []struct {
		input    []int32
		expected int32
	}{
		{[]int32{-2, 1, 3, -4, 5}, 8},
		{[]int32{3, 7, 4, 6, 5}, 13},
		{[]int32{2, 1, 5, 8, 4}, 11},
		{[]int32{3, 5, -7, 8, 10}, 15},
		{[]int32{-5, -5, -5, 10, -5}, 5},
	}

	for _, test := range tests {
		result := maxSubsetSum(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
