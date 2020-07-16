package main

// Пользователь даёт список заказа, программа должна по map
//с наименованиями товаров и ценами, посчитать сумму заказа.
//И сделать метод добавления новых товаров в map, и метод обновления цены уже существующего товара

import "fmt"

func addInShopList() func(item string, value float64) map[string]float64 {
	masBuy := make(map[string]float64)
	return func(item string, value float64) map[string]float64 {
		if item != "" {
			_, ok := masBuy[item]
			if ok {
				fmt.Printf("This item '%s' is already in the list, update it item price = %.1f\n", item, value)
			}
			masBuy[item] = value
		}
		return masBuy
	}
}

func sumPrice(masBuy func(item string, value float64) map[string]float64) string {
	var count float64
	for _, v := range masBuy("", 0) {
		count += v
	}
	return fmt.Sprintf("%.1f", count)
}

func getListBuys(masBuy func(item string, value float64) map[string]float64) map[string]float64 {
	return masBuy("", 0)
}

func main() {
	list := addInShopList()
	list("milk", 1.2)
	list("chocolate", 1.1)
	list("bread", 0.3)
	fmt.Println(getListBuys(list))
	fmt.Println(sumPrice(list))
	list("milk", 3.2)
	fmt.Println(getListBuys(list))
	fmt.Println(sumPrice(list))
}
