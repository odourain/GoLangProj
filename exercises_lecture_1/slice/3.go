package main
// Добавить в начало slice число 5
import (
	"fmt"
)

func addItem(slice []int, item int) (tempSlice []int)  {
	tempSlice = append(tempSlice, item)
	tempSlice = append(tempSlice, slice[:]...)
	return tempSlice
}

func main() {
	slice := [] int {1, 2, 3, 4}
	slice = addItem(slice, 5)
	fmt.Println(slice)
}
