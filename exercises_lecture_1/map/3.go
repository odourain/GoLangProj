package main
// Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих
import (
	"fmt"
)

func findInt(mas1 []int, mas2 []int) (mas []int) {
	m := make(map[int]bool)
	for i:=0; i<len(mas1); i++ {
		for j:=0; j<len(mas2); j++ {
			if _, ok := m[mas2[j]]; !ok && mas1[i] == mas2[j] {
				m[mas2[j]] = true
				mas = append(mas, mas2[j])
				i++
				j = -1
			}
		}
	}
	return mas
}

func main()  {
	m1:= []int {1, 2, 3, 5, 7, 8, 9, 11, 34, 42, 17, 54,33, 34}
	m2:= []int {1, 2, 4, 2, 11, 12, 4, 15, 34, 77, 27, 54,33, 34}
	fmt.Println(m1, m2)
	fmt.Println(findInt(m1, m2))
}