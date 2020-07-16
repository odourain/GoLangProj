package main

import (
	"fmt"
)

func findInt(mas1 []int, mas2 []int) []int {
	m := make(map[int]bool)
	var mas []int
	for _, v := range mas1 {
		for _, n := range mas2 {
			_, ok := m[n]
			if !ok && v == n{
				m[v] = true
				mas = append(mas, n)
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