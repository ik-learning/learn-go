package algs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func birthday(s []int32, d int32, m int32) int32 {
	result := 0
    for i:=0;i<len(s);i++{
    	if i+int(m) > len(s) {
    		break
	    }
    	count := d - s[i]
	    for z:=i+1;z<i+int(m);z++{
    		count -= s[z]
    		if count < 0 {
    			break
		    }
	    }
	    if count == 0 {
	    	result++
	    }
    }
	return int32(result)
}

func TestBirthDay(t *testing.T) {
	fixtures := [...]struct {
		chocolate []int32
		sum       int32
		squares   int32
		expected  int32
	}{
		{ []int32{1, 2, 1, 3, 2}, 3, 2, 2},
		{ []int32{1,1,1,1,1,1}, 3, 2, 0},
		{ []int32{4}, 4, 1, 1},
	}
	for _, e := range fixtures {
		result := birthday(e.chocolate, e.sum, e.squares)
		assert.Equal(t, e.expected, result, "they should be equal")
	}
}