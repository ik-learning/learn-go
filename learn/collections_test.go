package main

import (
	"strings"
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func TestCollectionV1(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}
	assert.Equal(t, 3, Index(strs, "plum"), "they should be equal")
	assert.Equal(t, false, Include(strs, "grape"), "they should be equal")
	assert.Equal(t, []string{"apple", "plum"}, Filter(strs, func(v string) bool {
		return strings.Contains(v, "l")
	}), "they should be equal")
	assert.Equal(t, []string{"PEACH", "APPLE", "PEAR", "PLUM"}, Map(strs, strings.ToUpper), "they should be equal")
}
