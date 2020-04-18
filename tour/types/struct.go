package main

import "fmt"

var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{} // values are set to 0
	p = &Vertex{1,2} // has type Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})
	StructFields()
	PointerToStruct()
	fmt.Println(v1, v2, v3, p)
}

type Vertex struct {
	X int
	Y int
}

func StructFields() {
	v := Vertex{X: 3, Y: 23}
	v.X = 4
	fmt.Println(v)
}

func PointerToStruct() {
	v := Vertex{X: 3, Y: 23}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
