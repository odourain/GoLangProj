package main
// Сдвинуть все элементы slice на 1 влево. Нулевой
// становится последним, первый - нулевым, последний - предпоследним.
import "fmt"

func shift(slice []int) []int {
	tempSlice := make([]int, len(slice))
	for i,v := range slice {
		if i-1 < 0 {
			tempSlice[len(tempSlice)-(-i+1)] = v
		} else {
			tempSlice[i-1] = v
		}
	}
	return tempSlice
}

func main()  {
	slice := []int {5, 1 ,2 ,3, 4}
	fmt.Println(slice)
	fmt.Println(shift(slice))
}
