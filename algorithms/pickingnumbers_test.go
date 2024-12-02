package algs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func pickingNumbers(a []int32) int32 {
	m := make(map[int32]int32)
	for i := int32(0); i < int32(len(a)); i++ {
		if _, exist := m[a[i]]; exist {
			m[a[i]]++
		} else {
			m[a[i]] = 1
		}
	}
	max := int32(2)
	// calculate pairs
	for key, count := range m {
		if max <= count {
			max = count
		}
		// check +1
		if value, exist := m[key+1]; exist {
			val := count + value
			if max <= val {
				max = val
			}
		}
		if value, exist := m[key-1]; exist {
			val := count + value
			if max <= val {
				max = val
			}
		}
	}
	return max
}

func TestShouldPickNumber(t *testing.T) {
	values := ReadFixturesInt("./utils/fixtures/picking-numbers.json")
	for i := 0; i < len(values); i++ {
		result := pickingNumbers(values[i].Input)
		fmt.Println(values[i].Input, result)
		assert.Equal(t, result, values[i].Result, "they should be equal")
	}
}

// fmt.Printf("Game > %d.", i)
// fmt.Printf(" Score > %d.", scores[i])
// fmt.Printf(" Highest Score > %d.", max)
// fmt.Printf(" Lowest Score > %d.\n", min)
