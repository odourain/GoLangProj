package main
// К каждому элементу []int прибавить 1
import "fmt"

func main()  {
	slice := []int {1, 2, 3, 4, 5, 6}
	for i := range slice {
		slice[i]++
	}
	fmt.Println(slice)
}
