package main
// Тоже, но сдвиг на i
import "fmt"

func shiftCustomRight(slice []int, shift int) []int {
	tempSlice := make([]int, len(slice))
	if shift > len(slice) || shift < 0 {
		fmt.Println("Your shift > len slice or < 0")
		return nil
	}
	for i,v := range slice {
		if i + shift > len(slice)-1 {
			tempSlice[i+shift-len(slice)] = v
		} else {
			tempSlice[i+shift] = v
		}
	}
	return tempSlice
}

func main()  {
	slice := []int {1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(shiftCustomRight(slice, 3))
}
