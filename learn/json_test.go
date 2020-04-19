package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type responseV2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func TestJsonV1(t *testing.T) {
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	assert.Equal(t, slcB, slcB, "they should be equal")

	res2D := &responseV2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	assert.Equal(t, res2B, res2B, "they should be equal")
}
