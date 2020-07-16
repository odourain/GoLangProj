package main

import "fmt"

func main(){
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice[len(slice)-1])
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
