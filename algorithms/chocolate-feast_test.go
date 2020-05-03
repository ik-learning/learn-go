package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func chocolateFeast(n int32, c int32, m int32) int32 {
	totalBars := n / c
	wrappers := n / c
	for  {
		if wrappers == 0 || wrappers < m {
			break
		}
		if wrappers != 0 && wrappers >= m  {
			totalBars += wrappers / m
			wrappers = wrappers + (wrappers / m) * (1 - m)
		}
	}
	return totalBars
}

func TestChocoFeast(t *testing.T) {
	fixtures := [...]struct {
		input []int32
		expected int32
	}{
		{[]int32{15, 3, 2}, 9},
		{[]int32{12, 4, 4}, 3},
		{[]int32{6, 2, 2}, 5},
	}
	for _, e := range fixtures {
		result := chocolateFeast(e.input[0], e.input[1], e.input[2])
		assert.Equal(t, e.expected, result, "they should be equal")
	}
}