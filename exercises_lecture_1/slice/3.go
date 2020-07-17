package main
// Добавить в начало slice число 5
import (
	"fmt"
)

func addItem(slice []int, item int) (tempSlice []int)  {
	for i, v := range slice{
		if i == 0 {
			tempSlice = append(tempSlice, item)
		}
		tempSlice = append(tempSlice, v)
	}
	return tempSlice
}

func main() {
	slice := [] int {1, 2, 3, 4}
	fmt.Println(slice)
	slice = addItem(slice, 5)
	fmt.Println(slice)
}
