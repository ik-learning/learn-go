package algs

import (
	"math"
	"testing"

	. "github.com/stretchr/testify/assert"
)

func appendAndDelete(s string, t string, k int32) string {
	if s == t {
		if k%2 == 0 && k > 1{
			return "Yes"
		} else if int32(len(s)/2) < k && int32(len(s)) != k{
			return "Yes"
		}
	}
	common := ""
	for i:=0;i<len(s) && i < len(t);i++{
		if s[i] == t[i] {
			common += string(s[i])
		} else {
			break
		}
	}
	if len(common) == 0 {
		full := int32(len(s) + len(t))
		if full <= k {
			return "Yes"
		}
	}
    ops := int32(math.Abs(float64((len(common) - len(s))+(len(common)-len(t)))))
    if ops == k {
    	return "Yes"
    } else if (ops - k) % 2 == 0 && (ops -k) < 0 {
	    return "Yes"
    }
    // perform full deletion
    if k > int32(len(s) + len(t)) {
    	return "Yes"
    }
	return "No"
}

func TestShouldAppendAndDelete(t *testing.T)  {
	fixtures := [...]struct {
		initial string
		desired string
		steps int32
		expected string
	}{
		{"aba", "aba", 2, "Yes"},
		{"aba", "aba", 7, "Yes"},
		{"aba", "aba", 3, "No"},
		{"hackerhappy", "hackerrank", 9, "Yes"},
		{"ashley", "ash", 2, "No"},
		{"aaaaaaaaaa", "aaaaa", 7, "Yes"},
		{"qwerasdf", "qwerbsdf", 6, "No"},
		{"abcdef", "fedcba", 15, "Yes"},
		{"aaa", "a", 5, "Yes"},
	}
	for _, e := range fixtures {
		result := appendAndDelete(e.initial, e.desired, e.steps)
		Equal(t, e.expected, result, "they should be equal")
	}
}

