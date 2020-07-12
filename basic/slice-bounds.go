package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 6, 11, 23}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
