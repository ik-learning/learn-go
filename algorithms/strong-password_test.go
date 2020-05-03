package algs

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minimumNumber(n int32, password string) int32 {
	added := 0
	isSingleNumber, _ := regexp.MatchString(`.*\d.*`, password)
	if !isSingleNumber {
		added += 1
	}
	isLowerCase, _ := regexp.MatchString(`.*[a-z].*`, password)
	if !isLowerCase {
		added += 1
	}
	isUpperCase, _ := regexp.MatchString(`.*[A-Z].*`, password)
	if !isUpperCase {
		added += 1
	}
	isSpecialChar, _ := regexp.MatchString(`.*[!@#$%^&*()\-+].*`, password)
	if !isSpecialChar {
		added += 1
	}
	if len(password) < 6 || len(password) + added < 6  {
		if len(password) + added < 6 {
			added += 6 - len(password) - added
		}
	}
	return int32(added)
}

func TestMinimumNumber(t *testing.T) {
	fixtures := [...]struct {
		pwd string
		expected  int32
	}{
		// {"Ab1", 3},
		// {"#HackerRank", 1},
		// {"4700", 3},
		{"AUzs-nV", 1},
		{"AUzs+nV", 1},
		{"AUzs+nV", 1},
	}
	for _, e := range fixtures {
		result := minimumNumber(int32(len(e.pwd)), e.pwd)
		assert.Equal(t, e.expected, result, "they should be equal")
	}
}
