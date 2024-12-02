package algs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// from 1700 to 1917 Julian Calendar
// from 1918 transition happens in January 31 > February 14
// from 1919 Gregorian Calendar
// February 29 leap year, 28 during all other year
func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func dayOfProgrammer(year int32) string {
	return ""
}

func TestShouldFindDayOfProgrammer(t *testing.T) {
	fixtures := [...]struct {
		input    int32
		expected string
	}{
		{2016, "12.09.2016"},
	}
	for _, e := range fixtures {
		result := dayOfProgrammer(e.input)
		Equal(t, e.expected, result, "they should be equal")
	}
}
