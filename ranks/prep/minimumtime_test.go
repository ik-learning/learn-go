package main

import (
	"fmt"
	"math"
	"testing"

	assert "github.com/stretchr/testify/assert"
)
// https://www.hackerrank.com/challenges/minimum-time-required/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=search
func minTime(machines []int64, goal int64) int64 {
	// How to calculate closed days?
	inday := float64(0)

	for _, val := range machines {
		inday += float64(1) / float64(val)
	}
	day := int64(float64(goal) / inday)
	fmt.Println(day)
	for {
		count := int64(0)
		for j := 0; j < len(machines); j++ {
			count += day / machines[j]
			if goal <= count {
				return day
			}
		}
		result := goal - count
		fmt.Println("Result >> ", result, day)
		day += int64(math.Floor(float64(result) / inday))
		fmt.Println("Result >> ", result, day)
		return -1
	}
}

func TestMinimumTime(t *testing.T) {
	var tests = []struct {
		machines []int64
		goal     int64
		expected int64
	}{
		//	{[]int64{2, 3, 2}, 10, 8},
		//{[]int64{2, 3}, 5, 6},
		{[]int64{4, 5, 6}, 12, 20},
		//	{[]int64{1, 3, 4}, 10, 7},
		//	{[]int64{63, 2, 26, 59, 16, 55, 99, 21, 98, 65}, 56, 82},
	}

	for _, test := range tests {
		result := minTime(test.machines, test.goal)
		assert.Equal(t, test.expected, result, "WTF")
	}
}
