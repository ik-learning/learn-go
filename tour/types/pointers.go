package main

import "fmt"

func main() {
	DoPointersV1()
}

func DoPointersV1() {
	i, j := 42, 213423

	p := &i // point to i
	fmt.Println(*p)
	*p = 23
	fmt.Println(i)

	p = &j
	*p = *p / 37

	fmt.Println(j)

}