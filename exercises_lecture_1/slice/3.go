package main

import (
	"fmt"
)

func addItem(slice []int, location int, item int) []int  {
	var tempSlice []int
	if location < 0 || location > len(slice) {
		fmt.Println("Location > size slice!")
		return tempSlice
	}
	if location == len(slice) {
		tempSlice = append(slice, item)
	} else {
		for i, v := range slice {
			if i == location {
				tempSlice = append(tempSlice, item)
				tempSlice = append(tempSlice, v)
			} else {
				tempSlice = append(tempSlice, v)
			}
		}
	}
	return tempSlice
}

func main() {
	slice := [] int {1, 2, 3, 4}
	fmt.Println(slice)
	slice = addItem(slice, 0, 5)
	fmt.Println(slice)
}
