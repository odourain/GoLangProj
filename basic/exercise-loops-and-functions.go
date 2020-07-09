package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (y float64) {
	z := float64(1)
	for {
		y = z-((math.Pow(z, 2)-x)/2*z)
		if int(y/0.001) == int(z/0.001) {
			return
		}
		z += 0.00001
	}
}

func main()  {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
