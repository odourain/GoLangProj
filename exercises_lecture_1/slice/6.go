package main

import "fmt"

func printAndDeleteItem(slice []int, location int)  []int{
	var tempSlice [] int
	if location < 0 || location >= len(slice) {
		fmt.Println("Location < 0 or > size slice")
		return tempSlice
	}
	fmt.Println(slice[location])
	for i, v:= range slice{
		if i != location {
			tempSlice = append(tempSlice, v)
		}
	}
	return tempSlice
}

func main()  {
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = printAndDeleteItem(slice, 4)
	fmt.Println(slice)
}
