package algs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func pageCount(n int32, p int32) int32 {
	if n == p {
		return 0
	}
	if p > n / 2 {
		if n % 2 == 0 {
			return n / 2 - p / 2
		} else {
			return (n - p) / 2
		}
	} else {
		turnPages := p / 2
		return turnPages
	}
}

type DrawingBook struct {
	pages int32
	page int32
	expect int32
}

func TestShouldCountPages(t *testing.T) {
	values := [...]DrawingBook{
		{6,2, 1},
		{5,4, 0},
		{99,12, 6},
		{6,5, 1},
		{ 4, 4, 0},
		{ 6 , 4, 1},
		{83246,78132, 2557},
	}
	for i := 0; i < len(values); i++ {
		result := pageCount(values[i].pages, values[i].page)
		fmt.Println(values[i].pages, values[i].page, result, values[i].expect)
		assert.Equal(t, result, values[i].expect, "they should be equal")
	}
}
