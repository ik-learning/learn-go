package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "./utils"
)

func birthdayCakeCandles(ar []int32) int32 {
	m := make(map[int32]int32)
	age := int32(0)
	for i := 0; i < len(ar); i++ {
		val := ar[i]
		if _, exist := m[val]; exist {
			m[val]++
		} else {
			m[val] = 1
			if age <= val {
				age = val
			}
		}
	}
	return m[age]
}

func TestShouldReturnFirstElement(t *testing.T) {
	read := ReadFixturesInt("./utils/fixtures/candles.json")

	values := read.Test
	for i := 0; i < len(values); i++ {
		result := birthdayCakeCandles(values[i].Input)
		assert.Equal(t, result, values[i].Result, "they should be equal")
	}
}
