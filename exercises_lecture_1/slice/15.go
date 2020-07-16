package main

import (
	"fmt"
)

func directOrder(slice []int) []int {
	check := false
	for i:=0; i< len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			temp:= slice[i]
			slice[i] = slice[i+1]
			slice[i+1] = temp
			check = true
		}
		if check && i+1 == len(slice)-1 {
			check = false
			i=-1
		}
	}
	return slice
}

func reverseOrder(slice []int) []int {
	check := false
	for i:=0; i< len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			temp:= slice[i]
			slice[i] = slice[i+1]
			slice[i+1] = temp
			check = true
		}
		if check && i+1 == len(slice)-1 {
			check = false
			i=-1
		}
	}
	return slice
}

func abcSort(slice []string) []string {
	check := false
	for i:=0; i< len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			temp:= slice[i]
			slice[i] = slice[i+1]
			slice[i+1] = temp
			check = true
		}
		if check && i+1 == len(slice)-1 {
			check = false
			i=-1
		}
	}
	return slice
}

func main()  {
	slice := []int {3, 4, 7, 1, 2, 7, 8, 9, 0}
	sliceABC := []string {"g","b","s","c","j","y","r","w","s","c","k"}
	fmt.Println(slice)
	fmt.Println(directOrder(slice))
	fmt.Println(reverseOrder(slice))
	fmt.Println(sliceABC)
	fmt.Println(abcSort(sliceABC))
}
