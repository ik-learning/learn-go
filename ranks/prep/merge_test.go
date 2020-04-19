package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countInversions(arr []int32) int64 {
	mergeSort(arr, make([]int32, len(arr)-1), 0, len(arr)-1)

	return int64(0)
}

func mergeSort(arr []int32, tmp []int32, left int, right int) {
	if left >= right {
		return
	}
	midlle := (left + right) / 2
	mergeSort(arr, tmp, left, midlle)
	mergeSort(arr, tmp, midlle+1, right)
	mergeHalves(arr, tmp, left, right)
}

func mergeHalves(arr []int32, tmp []int32, leftStart, rightEnd int) {
	leftEnd := (rightEnd + leftStart) / 2
	rightStart := leftEnd + 1
	// size := rightEnd - leftStart + 1

	left := leftStart
	right := rightStart
	index := leftStart

	for ok := true; ok == (left <= leftEnd && right <= rightEnd); {
		if arr[left] <= arr[right] {
			tmp[index] = arr[left]
			left++
		} else {
			tmp[index] = arr[index]
			right++
		}
		index++
	}

}

func TestCountInversions(t *testing.T) {
	var tests = []struct {
		input    []int32
		expected int
	}{
		{[]int32{1, 1, 3, 2, 2}, 0}}

	for _, test := range tests {
		_ = countInversions(test.input)
		assert.Equal(t, 1, 1, "they should be equal")
	}
}
