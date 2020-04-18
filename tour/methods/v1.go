package main

import (
	"fmt"
)

func main() {
	opV1()
	scaling()
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X*v.X - v.Y*v.Y
}

func Abs(v Vertex) float64 {
	return v.X*v.X - v.Y*v.Y
}

func opV1() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X + f
	v.Y = v.Y + f
}

func scaling() {
	v := Vertex{3, 6}
	v.Scale(10)
	fmt.Println(v)
}
