package main
// Взять последнее число slice, вернуть его пользователю, а из slice этот элемент удалить
import "fmt"

func main(){
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice[len(slice)-1])
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
