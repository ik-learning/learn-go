package main

import (
	"fmt"
)

func main() {
	fmt.Println(simple(20))
	fmt.Println(drop(23))
}

func simple(size int) int {
	sum := 0
	for i := 0; i < size; i++ {
		sum += 1
	}
	return sum
}

func drop(size int) int {
	sum := 0
	for sum < size {
		sum += 1
	}
	return sum
}