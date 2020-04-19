package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//https://www.hackerrank.com/challenges/grading/problem
func gradingStudents(grades []int32) []int32 {
	result := make([]int32, len(grades))
	for i, val := range grades {
		if val%5 == 0 || val < 38 {
			result[i] = val
		} else if val%5 < 3 {
			result[i] = val
		} else {
			result[i] = val + 5 - (val % 5)
		}
	}
	return result
}

func TestGradingStudents(t *testing.T) {
	var tests = []struct {
		input    []int32
		expected []int32
	}{
		{[]int32{73, 67, 38, 33}, []int32{75, 67, 40, 33}},
	}

	for _, test := range tests {
		result := gradingStudents(test.input)
		assert.Equal(t, test.expected, result, "they should be equal")
	}
}
