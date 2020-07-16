package main

import (
	"fmt"
)

func getCountWord(s string) (i int) {
	i++ // add last word
	for _, j := range s {
		if string(j) == " "{
			i++
		}
	}
	return i
}

func formatToStringsList(s string) []string {
	stringList:= make([]string, getCountWord(s))
	itemStringList:= 0
	for _, j := range s {
		if string(j) == " "{
			itemStringList++
		} else {
			stringList[itemStringList] += string(j)
		}
	}
	return stringList
}

func WordCount(s string) map[string] int {
	m:= make(map[string]int)
	count:= 0
	for i:=0;i<len(formatToStringsList(s));i++ {
		for j:=0;j<len(formatToStringsList(s));j++{
			if formatToStringsList(s)[i] == formatToStringsList(s)[j] {
				count++
			}
			if j == len(formatToStringsList(s))-1{
				m[formatToStringsList(s)[i]] = count
				count = 0
			}
		}
	}
	fmt.Println(s)
	fmt.Println(m)
	return map[string]int{"x" :1}
}

func main()  {
	WordCount("I am learning Go!")
}
