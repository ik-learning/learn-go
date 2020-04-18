package main

import "fmt"

func main() {
	ArrayV1()
	SliceV1()
	ArrSliceV1()
}

func ArrayV1() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(len(a), cap(a))

	primes := [6]int{2, 3, 5, 7}
	fmt.Println(primes, len(primes), cap(primes))
}

func SliceV1() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func ArrSliceV1() {
	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// len() and cap()
