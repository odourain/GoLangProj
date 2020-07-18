package main
// Тоже, что 9, но сдвиг вправо
import "fmt"

func shiftRight(slice []int) []int {
	slice = append(slice[len(slice)-1:], slice[:len(slice)-1]...)
	return slice
}

func main()  {
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(shiftRight(slice))
}
