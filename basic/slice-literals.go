package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 6, 11, 23}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{6, true},
		{11, false},
		{22, true},
	}
	fmt.Println(s)
}
