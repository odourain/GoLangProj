package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 6, 11, 85}
	fmt.Println(primes)
}
