package main
// Тоже, но сдвиг на заданное пользователем i
import "fmt"

func shiftCustom(slice []int, shift int) []int {
	tempSlice := make([]int, len(slice))
	if shift > len(slice) || shift < 0 {
		fmt.Println("Your shift > len slice or < 0")
		return nil
	}
	for i,v := range slice {
		if i - shift < 0 {
			tempSlice[len(tempSlice)-(-i+shift)] = v
		} else {
			tempSlice[i-shift] = v
		}
	}
	return tempSlice
}

func main()  {
	slice := []int {1, 2 ,3 ,4 ,5}
	fmt.Println(slice)
	fmt.Println(shiftCustom(slice, 2))
}
