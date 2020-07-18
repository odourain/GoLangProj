package main
// Есть текст, надо посчитать сколько раз каждое слова встречается.
import (
	"fmt"
	"strings"
)

func countWords(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		m[v]++
	}
	return m
}

func main()  {
	mapString := "Hello 5 world 4 and i am i sleep in i 5 a.m"
	fmt.Println(mapString)
	fmt.Println(countWords(mapString))
}
