package main

import "fmt"

func shiftCustom(slice []int, shift int) []int {
	tempSlice := make([]int, len(slice))
	if shift > len(slice) {
		fmt.Println("Your shift > len slice")
		return tempSlice
	}
	for i,v := range slice {
		if i - shift < 0 {
			tempSlice[len(tempSlice)-(-i+shift)] = v
		} else {
			tempSlice[i-shift] = v
		}
	}
	return tempSlice
}

func main()  {
	slice := []int {1, 2 ,3 ,4 ,5}
	fmt.Println(slice)
	fmt.Println(shiftCustom(slice, 2))
}
