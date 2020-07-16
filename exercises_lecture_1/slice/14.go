package main

import "fmt"

func swap(slice []int) []int {
	for i :=0; i<len(slice); i++ {
		if i+1 < len(slice) {
			temp := slice[i]
			slice[i] = slice[i+1]
			slice[i+1] = temp
			i++
		}
	}
	return slice
}

func main()  {
	slice := []int {0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice)
	fmt.Println(swap(slice))
}
