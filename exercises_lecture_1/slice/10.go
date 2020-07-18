package main
// Тоже, но сдвиг на заданное пользователем i
import "fmt"

func shiftCustom(slice []int, shift int) []int {
	if shift < 0 || shift > len(slice) {
		fmt.Println("This value < 0 or > len slice")
		return nil
	}
	slice = append(slice[shift:], slice[:shift]...)
	return slice
}

func main()  {
	slice := []int {1, 2 ,3 ,4 ,5}
	fmt.Println(slice)
	fmt.Println(shiftCustom(slice, 3))
}
