package main

import "fmt"

func main() {
	literalsV1()
	slicer()
}

func slicer() {
	s := []struct {
		i int
		b bool
	} {
		{2, true},
		{3, true},
		{7, false},
	}
	fmt.Println(s)
}


func literalsV1() {
	q := []int{2, 3, 5, 7, 11, 13}
	printSlice(q)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
