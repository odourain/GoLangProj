package main

import "fmt"

func main()  {
	slice := []int {1, 2 ,3 ,4}
	fmt.Println(slice)
	slice = append(slice, 5)
	fmt.Println(slice)
}
