package tdd1

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

// func BenchmarkCalc(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         _ = Calculate(i)
//     }
// }

var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

func TestCalculateMul(t *testing.T) {
	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Errorf("Test Failed: %d inputted, %d expected, received: %d", test.input, test.expected, output)
		}
	}
}
func TestCalculateV1(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 0},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		assert.Equal(Calculate(test.input), test.expected, "They should be equal")
	}
}