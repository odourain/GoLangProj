package main

import "fmt"

func splitSlice(slice1 [] int, slice2 [] int) []int {
	tempSlice := slice1
	for _,v := range slice2 {
		tempSlice = append(tempSlice, v)
	}
	return tempSlice
}

func main()  {
	slice1 := []int {1, 2, 3 ,4}
	slice2 := []int {5, 6, 7, 8}
	fmt.Println(slice1, slice2)
	fmt.Println(splitSlice(slice1, slice2))
}
