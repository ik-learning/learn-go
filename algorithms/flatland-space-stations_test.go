package algs

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func flatlandSpaceStations(n int32, c []int32) int32 {
	if len(c) == int(n) {
		return 0
	}
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	max := c[0]

	for i := 0; i < len(c); i++ {
		// check in between
		if i+1 < len(c) {
			current := c[i]
			next := c[i+1]
			// fmt.Println(current, next, (next - current) / 2)
			if max < (next -current)/2 {
				max = (next -current) / 2
			}
		} else {
			// check from right
			val := n - c[i] - 1
			if val > max {
				max = val
			}
		}
	}
	return max
}

func TestSpaceStation(t *testing.T) {
	fixtures := [...]struct {
		towns    int32
		stations []int32
		expected int32
	}{
		{20, []int32{13, 1, 11, 10, 6}, 6},
		{6, []int32{0, 1, 2, 3, 4, 5}, 0},
		{5, []int32{0, 4}, 2},
		{5, []int32{4}, 4},
		{5, []int32{0}, 4},
		{5, []int32{0, 3}, 1},
	}
	for _, e := range fixtures {
		result := flatlandSpaceStations(e.towns, e.stations)
		assert.Equal(t, e.expected, result, "they should be equal")
	}
}
