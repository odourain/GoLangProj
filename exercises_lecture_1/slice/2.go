package main
// Добавить в конец slice число 5
import "fmt"

func main()  {
	slice := []int {1, 2 ,3 ,4}
	slice = append(slice, 5)
	fmt.Println(slice)
}
