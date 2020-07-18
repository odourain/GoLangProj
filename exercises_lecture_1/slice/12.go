package main
// Тоже, но сдвиг на i
import "fmt"

func shiftCustomRight(slice []int, shift int) []int {
	if shift < 0 || shift > len(slice) {
		fmt.Println("This value < 0 or > len slice")
		return nil
	}
	slice = append(slice[len(slice)-shift:], slice[:len(slice)-shift]...)
	return slice
}

func main()  {
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(shiftCustomRight(slice, 1))
}
