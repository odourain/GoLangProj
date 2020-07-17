package main
// Взять первое число slice, вернуть его пользователю, а из slice этот элемент удалить
import "fmt"

func main()  {
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice[0])
	slice = slice[1:]
	fmt.Println(slice)
}
