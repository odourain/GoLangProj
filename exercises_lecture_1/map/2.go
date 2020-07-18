package main
// Есть очень большой массив или slice целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
import "fmt"

func findNumbers(mas []int) (masEnd []int) {
	m:= make(map[int]bool)
	for _, v := range mas {
		_, ok := m[v]
		if !ok {
			m[v] = true
			masEnd = append(masEnd, v)
		}
	}
	return
}

func main()  {
	mas := []int {1, 3, 5, 2, 1, 3, 5, 7, 8, 1, 2, 1, 4, 7, 5, 8, 9, 3, 4, 5, 7, 2, 8, 9, 5, 3, 2, 5, 3, 4, 7}
	fmt.Println(mas)
	fmt.Println(findNumbers(mas))
}
