package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	X, Y float64
}

func (v Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main()  {
	v := Vertex1{3, 4}
	fmt.Println(v.Abs())
}