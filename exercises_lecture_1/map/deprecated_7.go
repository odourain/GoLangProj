package main

import "fmt"

// К 2 добавить, чтобы хранились пользовательские аккаунты
// со счетом типа "вася: 300р, петя: 30000000р". И перед
// оформлением заказа, но после его расчёта мы проверяли,
// а есть ли деньги у пользователя, и если есть, то списывали
// сумму заказа.

type G struct {
	m map[string]float64
	history float64
}
var mapList2 map[int]G
//var history2 map[int]float64
var userBalance map[string]float64

func checkMaps2(m1 map[string]float64, m2 map[string]float64) bool {
	if len(m1) == len(m2) {
		count := 0
		for i, v := range m1 {
			if m2[i] == v {
				count++
			}
		}
		if count == len(m1) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func sumList2(m map[string]float64) float64 {
	sum := float64(0)
	for _, n := range m {
		sum += n
	}
	return sum
}

func checkBalanceForBuy(userName string, sale float64) string {
	if userBalance[userName] - sale <= 0 {
		return fmt.Sprintf("Transaction not work, because balance < price")
	} else {
		userBalance[userName] -= sale
		return fmt.Sprintf("Balance: %.1f", userBalance[userName])
	}
}

func getPriceList2(m map[string]float64, buyer string) string {
	_, ok := userBalance[buyer]
	if ok {
		for i, v := range mapList2 {
			if checkMaps2(v.m, m) {
				fmt.Printf("cache: ")
				return checkBalanceForBuy(buyer, mapList2[i].history)
			}
		}
		mapList2[len(mapList2)] = G{
			m: m,
			history: sumList2(m),
		}
		//history2[len(mapList2)-1] = sumList2(m)
		return checkBalanceForBuy(buyer, mapList2[len(mapList2)-1].history)
	} else {
		return fmt.Sprintf("Buyer not found.")
	}
}

func main() {
	//history2 = make(map[int]float64)
	mapList2 = make(map[int]G)
	userBalance = make(map[string]float64)
	userBalance = map[string]float64 {
		"Egor": 20,
		"John": 10,
		"Frank": 4,
	}
	m := map[string]float64{
		"milk":      1.2,
		"chocolate": 1.1,
		"bread":     0.3,
	}
	fmt.Println(getPriceList2(m, "Egor"))
	m1 := map[string]float64{
		"milk":      2.2,
		"chocolate": 1.1,
		"bread":     3.3,
	}
	fmt.Println(getPriceList2(m1, "John"))
	m2 := map[string]float64{
		"milk":      2.2,
		"chocolate": 1.1,
		"bread":     5.3,
	}
	fmt.Println(getPriceList2(m2, "Frank"))
	m3 := map[string]float64{
		"milk":      6.2,
		"chocolate": 1.1,
		"bread":     3.3,
	}
	fmt.Println(getPriceList2(m3, "Egor"))
	m4 := map[string]float64{
		"milk":      1.2,
		"chocolate": 1.1,
		"bread":     0.3,
	}
	fmt.Println(getPriceList2(m4, "Georg"))
	m5 := map[string]float64{
		"milk":      6.2,
		"chocolate": 1.1,
		"bread":     3.3,
	}
	fmt.Println(getPriceList2(m5, "Frank"))
}
