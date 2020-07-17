package main
// Вернуть пользователю копию переданного slice
import "fmt"

func copySlice(slice []int) []int{
	tempSlice := slice
	// or only return tempSlice
	return tempSlice
}

func main()  {
	slice := []int {1, 2 ,3 , 4}
	fmt.Println(copySlice(slice))
}
