package main

import (
	"fmt"
)

var c, python, java bool
var i, j int = 1, 2

func main() {
	println(add(34, 23))
	println(subtract(98, 16))
	a, b := swap("hello", "world")
	println(a, b)
	fmt.Println(nakedReturn(78))
	variables()
	varsInit()
}

func add(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func nakedReturn(sum int) (x, y int) {
	x = sum * 7/9
	y = sum - x
	return
}

func variables() {
	var i int
	fmt.Println(i, c, python, java)
}

func varsInit() {
	var c, python, java = true, false, "no!"
	k := 78
	fmt.Println(i, j, c, python, java, k)
}