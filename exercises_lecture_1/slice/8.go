package main
// Из первого slice удалить все числа, которые есть во втором
import "fmt"

func deleteSliceOneWithSliceTwo(slice1 []int, slice2 []int) (tempSlice []int) {
	for _, v1 := range slice2 {
		for i, v2 := range slice1 {
			if v1 == v2 {
				slice1 = append(slice1[:i], slice1[i+1:]...)
			}
		}
	}
	return slice1
}

func main()  {
	slice1 := []int {1, 2, 3, 4, 5}
	slice2 := []int {2, 4, 6, 7}
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice1 = deleteSliceOneWithSliceTwo(slice1, slice2)
	fmt.Println(slice1)
}
