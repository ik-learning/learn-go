package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func hurdleRace(k int32, height []int32) int32 {
	result := k
	for i := 0; i < len(height); i++ {
		if result < height[i] {
			result = height[i]
		}
	}
	return result - k
}

func TestHurdleRace(t *testing.T) {
	fixtures := [...]struct {
		height   int32
		barriers []int32
		expected int32
	}{
		{4, []int32{1, 6, 3, 5, 2}, 2},
		{7, []int32{2, 5, 4, 5, 2}, 0},
		{53, []int32{86, 4, 83, 20, 6, 81, 58, 59, 53, 2, 54, 62, 25, 35, 79, 64, 27, 49, 32, 95, 100,
			20, 58, 39, 92, 30, 67, 89, 58, 81, 100, 66, 73, 29, 75, 81, 70, 55, 18, 28, 7, 35, 98, 52, 30, 11, 69, 48,
			84, 54, 13, 14, 15, 86, 34, 82, 92, 26, 8, 53, 62, 57, 50, 31, 61, 85, 88, 5, 80, 64, 90, 52, 47, 43, 40, 93,
			69, 70, 16, 43, 7, 25, 99, 12, 63, 99, 71, 76, 55, 17, 90, 43, 27, 20, 42, 84, 39, 96, 75, 1}, 47},
	}
	for _, e := range fixtures {
		result := hurdleRace(e.height, e.barriers)
		Equal(t, e.expected, result, "they should be equal")
	}
}
