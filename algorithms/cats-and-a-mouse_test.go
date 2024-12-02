package algs

import (
	. "math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func catAndMouse(x int32, y int32, z int32) string {
	// from cat A
	fromA := int32(Abs(float64(x - z)))
	fromB := int32(Abs(float64(y - z)))

	if fromA-fromB == 0 {
		return "Mouse C"
	} else if fromA > fromB {
		return "Cat B"
	} else {
		return "Cat A"
	}
}

func TestShouldFindMouse(t *testing.T) {
	fixtures := [...]struct {
		catA     int32
		catB     int32
		mouseC   int32
		expected string
	}{
		{1, 2, 3, "Cat B"},
		{1, 3, 2, "Mouse C"},
	}
	for _, e := range fixtures {
		result := catAndMouse(e.catA, e.catB, e.mouseC)
		Equal(t, e.expected, result, "they should be equal")
	}
}
