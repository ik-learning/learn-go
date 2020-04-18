package main

import (
	"testing"
	"time"
	// assert "github.com/stretchr/testify/assert"
)

func worker(values []string, done chan string) {
	time.Sleep(time.Second)
	done <- values[0]

}

func TestThreadingSyncV1(t *testing.T) {
	var strs = []string{"peach", "apple", "pear", "plum"}

	done := make(chan string, 1)
	go worker(strs, done)
    <-done
	//fmt.Println(<-done)
}
