package main

import "fmt"

func countWords(s string) map[string]int {
	m := make(map[string]int)
	tempString := ""
	for i, v := range s {
		if string(v) == " " && i != 0{
			_, ok := m[tempString]
			if ok {
				m[tempString]++
			} else {
				m[tempString] = 1
			}
			tempString = ""
		} else {
			tempString += string(v)
		}
	}
	return m
}

func main()  {
	mapString := "Hello 5 world 4 and i am i sleep in i 5 a.m"
	fmt.Println(mapString)
	fmt.Println(countWords(mapString))
}
