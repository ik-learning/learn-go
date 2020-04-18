package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(msg string) {
		time.Sleep(2 * time.Second)
		c1 <- msg
	}("result 1")

	go func(msg string) {
		time.Sleep(2 * time.Second)
		c2 <- msg
	}("result 2")

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }

}

// time go run threading/timeouts.go
