package main

import "fmt"

// Сделать 1е, но у нас приходит несколько сотен таких
// списков заказов и мы хотим запоминать уже посчитанные
// заказы, чтобы если встречается такой же, то сразу
// говорить его цену без расчёта.
var mapList map[int]historyBuy

type historyBuy struct {
	buyList map[string]float64
	history float64
}

//func checkMaps(m1 map[string]float64, m2 map[string]float64) bool {
//	if len(m1) == len(m2) {
//		count := 0
//		for i, v := range m1 {
//			if m2[i] == v {
//				count++
//			}
//		}
//		if count == len(m1) {
//			return true
//		} else {
//			return false
//		}
//	} else {
//		return false
//	}
//}

func sumList(m map[string]float64) float64 {
	sum := float64(0)
	for _, n := range m {
		sum += n
	}
	return sum
}

func getPriceList(m map[string]float64) string {
	for _, v := range mapList {
		if checkMaps(v.buyList, m) {
			fmt.Printf("cache: ")
			return fmt.Sprintf("%.1f", v.history)
		}
	}
	mapList[len(mapList)] = historyBuy{
		buyList: m,
		history: sumList(m),
	}
	return fmt.Sprintf("%.1f", mapList[len(mapList)-1].history)
}

func main() {
	mapList = make(map[int]historyBuy)
	m := map[string]float64{
		"milk":      1.2,
		"chocolate": 1.1,
		"bread":     0.3,
	}
	fmt.Println(getPriceList(m))
	m1 := map[string]float64{
		"milk":      2.2,
		"chocolate": 1.1,
		"bread":     3.3,
	}
	fmt.Println(getPriceList(m1))
	m2 := map[string]float64{
		"milk":      2.2,
		"chocolate": 1.1,
		"bread":     5.3,
	}
	fmt.Println(getPriceList(m2))
	m3 := map[string]float64{
		"milk":      6.2,
		"chocolate": 1.1,
		"bread":     3.3,
	}
	fmt.Println(getPriceList(m3))
	m4 := map[string]float64{
		"milk":      1.2,
		"chocolate": 1.1,
		"bread":     0.3,
	}
	fmt.Println(getPriceList(m4))
	m5 := map[string]float64{
		"milk":      6.2,
		"chocolate": 1.1,
		"bread":     3.3,
	}
	fmt.Println(getPriceList(m5))
}
