package main
// Вернуть пользователю копию переданного slice
import "fmt"

func copySlice(slice []int) []int{
	tempSlice := make([]int, len(slice))
	copy(tempSlice, slice)
	return tempSlice
}

func main()  {
	slice := []int {1, 2 ,3 , 4}
	fmt.Println(copySlice(slice))
}
