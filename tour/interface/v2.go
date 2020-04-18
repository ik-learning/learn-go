package main

import "fmt"

func main() {
	var i I = T{"hello"}
	i.M()
}

type I interface {
	M()
}

type T struct {
	text string
}

func (t T) M() {
	fmt.Println(t.text)
}

