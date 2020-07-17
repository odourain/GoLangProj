package main
// К каждому элементу []int прибавить 1
import "fmt"

func main()  {
	slice := []int {1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	for i:= 0; i< len(slice); i++ {
		slice[i]++
	}
	fmt.Println(slice)
}
