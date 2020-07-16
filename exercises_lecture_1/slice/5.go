package main

import "fmt"

func main()  {
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice[0])
	slice = slice[1:]
	fmt.Println(slice)
}
