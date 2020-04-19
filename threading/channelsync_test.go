package main

import (
	"testing"
	// assert "github.com/stretchr/testify/assert"
)

func TestThreadingV2(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}

	messages := make(chan string, 2)

	messages <- strs[1]
	messages <- strs[0]

	//fmt.Println(len(messages))
	//fmt.Println(<-messages)
}