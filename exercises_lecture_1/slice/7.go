package main
// Объединить два slice и вернуть новый со всеми элементами первого и второго
import "fmt"

func splitSlice(slice1 [] int, slice2 [] int) []int {
	for _,v := range slice2 {
		slice1 = append(slice1, v)
	}
	return slice1
}

func main()  {
	slice1 := []int {1, 2, 3 ,4}
	slice2 := []int {5, 6, 7, 8}
	fmt.Println(slice1, slice2)
	fmt.Println(splitSlice(slice1, slice2))
}
