package main
// Взять i-е число slice, вернуть его пользователю, а
// из slice этот элемент удалить. Число i передает пользователь в функцию
import "fmt"

func printAndDeleteItem(slice []int, location int) (tempSlice []int){
	if location < 0 || location >= len(slice) {
		fmt.Println("Location < 0 or > size slice")
		return nil
	}
	fmt.Println(slice[location])
	tempSlice = append(slice[:location], slice[location+1])
	return tempSlice
}

func main()  {
	slice := []int {1, 2, 3, 4, 5}
	slice = printAndDeleteItem(slice, 3)
	fmt.Println(slice)
}
