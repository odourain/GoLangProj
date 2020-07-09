package main

import (
	"fmt"
	"math"
)

const Pi = math.Pi

func main()  {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
