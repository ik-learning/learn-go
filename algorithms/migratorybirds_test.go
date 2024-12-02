package algs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func migratoryBirds(arr []int32) int32 {
	m := make(map[int32]int32)
	max := int32(1)
	for i := 0; i < len(arr); i++ {
		val := arr[i]
		if _, exist := m[val]; exist {
			m[val]++
			if max < m[val] {
				max = m[val]
			}
		} else {
			m[val] = 1
		}
	}
	min := int32(10)
	for key, element := range m {
		if element == max {
			if key < min {
				min = key
			}
		}
	}
	return min
}

func TestShouldProcessMigratoryBirds(t *testing.T) {
	values := ReadFixturesInt("./utils/fixtures//migratory-birds.json")
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i].Input)
		result := migratoryBirds(values[i].Input)
		assert.Equal(t, result, values[i].Result, "they should be equal")
	}
}
