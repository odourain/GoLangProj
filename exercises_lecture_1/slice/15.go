package main

// Упорядочить slice в порядке: прямом, обратном, лексикографическом.
import (
	"fmt"
	"sort"
)

func directOrder(slice []int) []int {
	sort.Ints(slice)
	return slice
}

func reverseOrder(slice []int) []int {
	tempSlice := make([]int, 0, len(slice))
	sort.Ints(slice)
	for i := len(slice) - 1; i >= 0; i-- {
		tempSlice = append(tempSlice, slice[i])
	}
	return tempSlice
}

func abcSort(slice []string) []string {
	sort.Strings(slice)
	return slice
}

func main() {
	slice := []int{3, 4, 7, 1, 2, 7, 8, 9, 0}
	sliceABC := []string{"g", "b", "s", "c", "j", "y", "r", "w", "s", "c", "k"}
	fmt.Println(slice)
	fmt.Println(directOrder(slice))
	fmt.Println(reverseOrder(slice))
	fmt.Println(sliceABC)
	fmt.Println(abcSort(sliceABC))
}
