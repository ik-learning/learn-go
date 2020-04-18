package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(msg string) {
		time.Sleep(1 * time.Second)
		c1 <- msg
	}("one")

	go func(msg string) {
		time.Sleep(2 * time.Second)
		c2 <- msg
	}("two")

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// time go run threading/select.go
