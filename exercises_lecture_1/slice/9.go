package main
// Сдвинуть все элементы slice на 1 влево. Нулевой
// становится последним, первый - нулевым, последний - предпоследним.
import "fmt"

func shift(slice []int) []int {
	slice = append(slice[1:], slice[:1]...)
	return slice
}

func main()  {
	slice := []int {5, 1 ,2 ,3, 4}
	fmt.Println(slice)
	fmt.Println(shift(slice))
}
