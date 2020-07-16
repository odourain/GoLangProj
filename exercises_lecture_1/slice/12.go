package main

import "fmt"

func shiftCustomRight(slice []int, shift int) []int {
	tempSlice := make([]int, len(slice))
	if shift > len(slice) {
		fmt.Println("Your shift > len slice")
		return tempSlice
	}
	for i,v := range slice {
		if i + shift > len(slice)-1 {
			tempSlice[i+shift-len(slice)] = v
		} else {
			tempSlice[i+shift] = v
		}
	}
	return tempSlice
}

func main()  {
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(shiftCustomRight(slice, 3))
}
