package algs

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	result := int64(-1)
	if bc == wc {
		result = int64(b*bc + w*wc)
	}
	if math.Abs(float64(bc - wc)) > float64(z) {
		if bc > wc {
			bc = wc + z
		} else {
			wc = bc + z
		}
		result = int64(b)*int64(bc) + int64(w)*int64(wc)
	} else {
		result =  int64(b)*int64(bc) + int64(w)*int64(wc)
	}
	return result
}

func TestShouldTaumBDay(t *testing.T) {
	for _, e := range fixtures {
		result := taumBday(e.black, e.white, e.bc, e.wc, e.z)
		assert.Equal(t, e.expected, result, "they should be equal")
	}
}

type testFixtures struct {
	black int32
	white int32
	bc int32
	wc int32
	z int32
	expected int64
}

var fixtures = []testFixtures {
	{10, 10, 1, 1, 1, 20},
	{5, 9, 2, 3, 4, 37},
	{3, 6, 9, 1, 1, 12},
	{7, 7, 4, 2, 1, 35},
	{3, 3, 1, 9, 2, 12},
	{3, 0, 6, 9, 2, 18},
	{27984, 1402, 619246, 615589, 247954, 18192035842},
	{95677, 39394, 86983, 311224, 588538, 20582630747},
	{68406, 12718, 532909, 315341, 201009, 39331944938},
	{15242, 95521, 712721, 628729, 396706, 70920116291},
	{21988, 67781, 645748, 353796, 333199, 38179353700},
	{22952, 80129, 502975, 175136, 340236, 25577754744},
	{38577, 3120, 541637, 657823, 735060, 22947138309},
	{5943, 69851, 674453, 392925, 381074, 31454478354},
	{62990, 61330, 310144, 312251, 93023, 38686324390},
	{11152, 43844, 788543, 223872, 972572, 18609275504},
}

