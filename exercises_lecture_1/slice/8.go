package main
// Из первого slice удалить все числа, которые есть во втором
import "fmt"

func deleteSliceOneWithSliceTwo(slice1 []int, slice2 []int) (tempSlice []int)  {
	deleteItem := false
	for _,v := range slice1 {
		for _,k := range slice2 {
			if v == k {
				deleteItem = true
			}
		}
		if !deleteItem{
			tempSlice = append(tempSlice, v)
		}
		deleteItem = false
	}
	return tempSlice
}

func main()  {
	slice1 := []int {1, 2, 3, 4, 5}
	slice2 := []int {2, 4, 6, 7}
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice1 = deleteSliceOneWithSliceTwo(slice1, slice2)
	fmt.Println(slice1)
}
