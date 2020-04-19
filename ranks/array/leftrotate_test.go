package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func leftRotate(arr []int32, rotate int32) []int32 {
	length := int32(len(arr))
	rotateRight := length - rotate
	tmpRight := arr[0 : length-rotateRight]
	tmpLeft := arr[length-rotateRight : length]
	return append(tmpLeft, tmpRight...)
}

func TestLeftRotate(t *testing.T) {
	var tests = []struct {
		input    []int32
		rotate   int32
		expected []int32
	}{
		{[]int32{1, 2, 3, 4, 5}, 4, []int32{5, 1, 2, 3, 4}},
		{[]int32{1, 2, 3, 4, 5}, 2, []int32{3, 4, 5, 1, 2}},
	}

	for _, test := range tests {
		result := leftRotate(test.input, test.rotate)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
