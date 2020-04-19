package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})
	var count int32 = 0
	for i := 0; i < len(prices); i++ {
		if k -= prices[i]; k > 0 {
			count += 1
		} else {
			break
		}
	}
	return count
}

func TestMaximumToys(t *testing.T) {
	var tests = []struct {
		prices   []int32
		sum      int32
		expected int32
	}{
		{[]int32{1, 4, 3, 2}, 7, 3},
		{[]int32{1, 4, 3, 2, 6, 8}, 20, 5},
		{[]int32{1, 12, 5, 111, 200, 1000, 10}, 50, 4},
	}

	for _, test := range tests {
		result := maximumToys(test.prices, test.sum)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
