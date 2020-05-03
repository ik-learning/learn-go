package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func permutationEquation(p []int32) []int32 {
	f1 := map[int32]int32{}
    result := make([]int32, len(p))
    for i := 0;i<len(p);i++{
    	f1[p[i]] = int32(i) + 1
    }
    for i := 0;i<len(p);i++{
    	result[i] = f1[f1[int32(i+1)]]
    }
	return result
}


func TestPermutationEquitation(t *testing.T) {
	fixtures := [...]struct {
		input    []int32
		expected []int32
	}{
		{[]int32{2,3,1}, []int32{2,3,1}},
		{[]int32{4, 3, 5, 1, 2}, []int32{1,3,5,4,2}},
	}
	for _, e := range fixtures {
		result := permutationEquation(e.input)
		assert.Equal(t, e.expected, result, "they should be equal")
	}
}
