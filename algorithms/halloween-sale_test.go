package algs

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

// Return the number of games you can buy
func howManyGames(p int32, d int32, m int32, s int32) int32 {
	if p > s {
		return int32(0)
	}
	credits := s - p
	count := int32(1)
	gameCost := p
	for ok:=true; ok; ok = credits - gameCost > 0 {
		if gameCost - d > m {
			gameCost -=d
			if credits - gameCost > 0 {
				credits -= gameCost
				count++
			}
		} else if gameCost == m && credits - m > 0 {
			count += credits / m
			break
		} else if gameCost - d <= m && credits - m > 0 {
			gameCost = m
			if credits - gameCost > 0 {
				credits -= gameCost
				count++
			}
		} else {
			break
		}
	}
	return count
}

func TestHowManyGamesToBuy(t *testing.T) {
	fixtures := [...]struct {
		input    []int32
		expected int32
	}{
		{[]int32{20, 3, 6, 80}, 6},
		{[]int32{20, 3, 6, 85}, 7},
		{[]int32{16, 2, 1, 9981}, 9917},
		{[]int32{100, 1, 1, 99}, 0},
		{[]int32{100, 1, 1, 100}, 1},
		{[]int32{100, 19, 1, 180}, 1},
	}
	for _, e := range fixtures {
		result := howManyGames(e.input[0], e.input[1], e.input[2], e.input[3])
		Equal(t, e.expected, result, "they should be equal")
	}
}
