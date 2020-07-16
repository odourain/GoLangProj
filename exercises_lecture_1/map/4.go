package main

// 4. Сделать Фибоначчи с мемоизацией (https://ru.wikipedia.org/wiki/%D0%9C%D0%B5%D0%BC%D0%BE%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D1%8F)
import "fmt"

func fib() func(int) []int {
	history := make(map[int]int)
	return func(a int) []int {
		var m []int
		_, ok := history[a-1]
		if ok {
			fmt.Printf("cache: ")
			for i := 0; i < a; i++ {
				m = append(m, history[i])
			}
			return m
		} else {
			for i := 0; i < a; i++ {
				switch {
				case i > 1:
					sum := history[i-1] + history[i-2]
					history[i] = sum
					m = append(m, sum)
				default:
					history[i] = i
					m = append(m, i)
				}
			}
			return m
		}
	}
}

func main() {
	f := fib()
	fmt.Println(f(10000))
	fmt.Println(f(9999))
}
