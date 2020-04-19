package algs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	. "./utils"
)

func breakingRecords(scores []int32) []int32 {
	max := scores[0]
	maxCount := int32(0)
	min := scores[0]
	minCount := int32(0)
	for i := 0; i < len(scores); i++ {
		if max < scores[i] {
			max = scores[i]
			maxCount++
		}
		if min > scores[i] {
			min = scores[i]
			minCount++
		}
	}
	return []int32{maxCount, minCount}
}

func TestShouldProcessBreakingRecords(t *testing.T) {
	values := ReadFixturesIntV2("./utils/fixtures/breaking-the-records.json")
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i].Input)
		result := breakingRecords(values[i].Input)
		fmt.Println(result)
		assert.Equal(t, result, values[i].Result, "they should be equal")
	}
}

// fmt.Printf("Game > %d.", i)
// fmt.Printf(" Score > %d.", scores[i])
// fmt.Printf(" Highest Score > %d.", max)
// fmt.Printf(" Lowest Score > %d.\n", min)