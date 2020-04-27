package algs

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func findDigits(n int32) int32  {
	inStr := fmt.Sprintf("%d", n)
	result := int32(0)

	for _, e := range inStr {
		val, err := strconv.ParseInt(fmt.Sprintf("%c", e), 10, 32)
		if err == nil &&  val != 0 && n%int32(val) == 0 {
			result += 1
		}
	}

	return result
}

func TestShouldFindDigits(t *testing.T) {
	fixtures := [...]struct {
		input int32
		expected int32
	}{
		{12, 2},
		{1012, 3},
		{701209803, 3},
	}
	for _, e := range fixtures {
		result := findDigits(e.input)
		Equal(t, e.expected, result, "they should be equal")
	}
}



