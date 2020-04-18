package main

import (
	"fmt"
)

func main() {
	iterator()
}

var pow = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func iterator() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}