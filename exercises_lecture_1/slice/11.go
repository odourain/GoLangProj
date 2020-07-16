package main

import "fmt"

func shiftRight(slice []int) []int {
	tempSlice := make([]int, len(slice))
	for i,v := range slice {
		if i + 1 > len(slice)-1 {
			tempSlice[len(slice)-(i+1)] = v
		} else {
			tempSlice[i+1] = v
		}
	}
	return tempSlice
}

func main()  {
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(shiftRight(slice))
}
