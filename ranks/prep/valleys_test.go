package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
		steps    int32
		path     string
		expected int
	}{
		{8, "UDDDUU", 1},
		{8, "DDUUUUDD", 1},
		{8, "UDDDUDUU", 1},
	}


func TestCountingValleys(t *testing.T) {
	for _, test := range tests {
		val := countingValleys(test.steps, test.path)
		assert.Equal(t, int32(test.expected), val, "they should be equal")
	}
}
