package algs

import (
	"math"
	"testing"
)

func serviceLane(lane []int32, cases [][]int32) []int32 {
	var result []int32

	for _, e := range cases {
		min := int32(math.MaxInt32)
		tmp := lane[e[0] : e[1]+1]
		for _, i := range tmp {
			if min > i {
				min = i
			}
		}
		result = append(result, min)
	}
	return result
}

func TestShouldCalculateServiceLane(t *testing.T) {
	fixtures := [...]struct {
		lane     []int32
		cases    [][]int32
		expected []int32
	}{
		{[]int32{2, 3, 1, 2, 3, 2, 3, 3}, [][]int32{
			{0, 3},
			{4, 6},
			{6, 7},
			{3, 5},
			{0, 7},
		}, []int32{1, 2, 3, 2, 1}},
		{[]int32{1, 2, 2, 2, 1}, [][]int32{
			{2, 3},
			{1, 4},
			{2, 4},
			{2, 4},
			{2, 3},
		}, []int32{2, 1, 1, 1, 2}},
	}
	for _, e := range fixtures {
		result := serviceLane(e.lane, e.cases)
		Equal(t, e.expected, result, "they should be equal")
	}
}
