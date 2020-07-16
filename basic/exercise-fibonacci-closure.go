package main

import "fmt"

func fibonacci() func(int) int {
	var s []int
	sum := 1
	return func(a int) int {
		if a >= 2 {
			sum = s[a-1] + s[a-2]
			s = append(s, sum)
			return sum
		} else {
			s = append(s, a)
			return a
		}
	}
}

func main()  {
	f:= fibonacci()
	for i:=0; i<10;i++ {
		fmt.Println(f(i))
	}
}
