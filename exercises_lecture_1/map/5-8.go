package main

// Пользователь даёт список заказа, программа должна по map
//с наименованиями товаров и ценами, посчитать сумму заказа.
//И сделать метод добавления новых товаров в map, и метод обновления цены уже существующего товара

// Сделать 1е, но у нас приходит несколько сотен таких
// списков заказов и мы хотим запоминать уже посчитанные
// заказы, чтобы если встречается такой же, то сразу
// говорить его цену без расчёта.

// К 2 добавить, чтобы хранились пользовательские аккаунты
// со счетом типа "вася: 300р, петя: 30000000р". И перед
// оформлением заказа, но после его расчёта мы проверяли,
// а есть ли деньги у пользователя, и если есть, то списывали
// сумму заказа.

// Есть map аккаунтов и счетов, как описано в 3.
// Надо вывести ее в отсортированном виде с сортировкой:
// по имени в алфавитном порядке, по имени в обратном
// порядке, по количеству денег по убыванию.

import (
	"fmt"
	"sort"
)

type shopListHistory struct {
	list    map[int]shopList
	sumList float64
}

type shopList struct {
	list    map[string]float64
	sumList float64
}

type clientListBuys struct {
	clientHistory shopListHistory
	balance       float64
}

type clientInformationList struct {
	name    string
	balance float64
}

func addInShopList() func(item string, value float64) map[string]float64 {
	masBuy := make(map[string]float64)
	return func(item string, value float64) map[string]float64 {
		if item != "" {
			masBuy[item] = value
		}
		return masBuy
	}
}

func addListShopsList() func(masBuy *func(item string, value float64) map[string]float64) shopListHistory {
	listBuys := shopListHistory{
		make(map[int]shopList),
		0,
	}
	return func(masBuy *func(item string, value float64) map[string]float64) shopListHistory {
		if masBuy == nil {
			return listBuys
		}
		for i := range listBuys.list {
			if checkMaps((*masBuy)("", 0), listBuys.list[i].list) {
				masBuy = nil
				return listBuys // cache
			}
		}
		listBuys.list[len(listBuys.list)] = shopList{
			(*masBuy)("", 0),
			0,
		}
		listBuys = sumPrice(listBuys, len(listBuys.list)-1)
		masBuy = nil
		return listBuys
	}
}

func addClientListBuys(clients []clientInformationList) func(string, shopListHistory) (map[string]clientListBuys, string, float64) {
	clientHistory := make(map[string]clientListBuys)
	for _, v := range clients {
		clientHistory[v.name] = clientListBuys{
			shopListHistory{
				make(map[int]shopList),
				0},
			v.balance,
		}
	}
	return func(nameClient string, buys shopListHistory) (map[string]clientListBuys, string, float64) {
		if nameClient == "" {
			return clientHistory, nameClient, 0
		}
		v, ok := clientHistory[nameClient]
		if !ok {
			fmt.Println("This client is not found. Add this client in database.")
			clientHistory[nameClient] = clientListBuys{
				shopListHistory{
					make(map[int]shopList),
					0},
				0,
			}
			fmt.Printf("Add new user '%s'.\n", nameClient)
			return clientHistory, nameClient, clientHistory[nameClient].balance
		}
		if v.balance-buys.sumList < 0 {
			fmt.Println("insufficient funds.")
		} else {
			clientHistory[nameClient] = clientListBuys{
				shopListHistory{
					buys.list,
					buys.sumList,
				},
				v.balance - buys.sumList,
			}
		}
		return clientHistory, nameClient, clientHistory[nameClient].balance
	}
}

func printBalance(_ map[string]clientListBuys, nameClient string, balance float64) string {
	return fmt.Sprintf("Balance client '%s' = %.1f", nameClient, balance)
}

func sortClientsName(m map[string]clientListBuys, _ string, _ float64) {
	clients := make([]string, 0, len(m))
	for i := range m {
		clients = append(clients, i)
	}
	sort.Strings(clients)
	fmt.Printf("\n")
	for _, v := range clients {
		fmt.Printf("Balance '%s' = %.1f \n", v, m[v].balance)
	}
}

func sortClientsNameReverse(m map[string]clientListBuys, _ string, _ float64) {
	clients := make([]string, 0, len(m))
	for i := range m {
		clients = append(clients, i)
	}
	sort.Strings(clients)
	fmt.Printf("\n")
	for i := len(clients) - 1; i >= 0; i-- {
		fmt.Printf("Reverse Balance '%s' = %.1f \n", clients[i], m[clients[i]].balance)
	}
}

func sortClientsBalanceReverse(m map[string]clientListBuys, _ string, _ float64) {
	client := make([]float64, 0, len(m))
	for _, v := range m {
		client = append(client, v.balance)
	}
	sort.Float64s(client)
	fmt.Printf("\n")
	for i := len(client) - 1; i >= 0; i-- {
		for j, v := range m {
			if client[i] == v.balance{
				fmt.Printf("Reverse names for balance '%s' = %.1f \n", j, client[i])
				delete(m, j)
			}
		}
	}
}

func sumPrice(listBuys shopListHistory, item int) shopListHistory {
	var count float64
	for _, v := range listBuys.list[item].list {
		count += v
	}
	listBuys.list[item] = shopList{
		listBuys.list[item].list,
		count,
	}
	listBuys.sumList += count
	return listBuys
}

func checkMaps(m1 map[string]float64, m2 map[string]float64) bool {
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

func main() {
	clients := []clientInformationList{
		{"Frank", 30},
		{"Putin", 999999999},
		{"Rich", 150},
	}

	mass := addInShopList()
	mass("milk", 1.2)
	mass("chocolate", 1.1)
	mass("bread", 0.3)

	mass2 := addInShopList()
	mass2("milk", 2.2)
	mass2("chocolate", 1.1)
	mass2("bread", 4.3)

	mass3 := addInShopList()
	mass3("milk", 1.2)
	mass3("chocolate", 1.1)
	mass3("bread", 0.3)

	clientList := addClientListBuys(clients)
	mass1 := addInShopList()
	mass1("milk", 13.2)
	mass1("chocolate", 1.1)
	mass1("bread", 6.3)

	mass21 := addInShopList()
	mass21("milk", 7.2)
	mass21("chocolate", 1.1)
	mass21("bread", 19.3)

	mass31 := addInShopList()
	mass31("milk", 12.2)
	mass31("chocolate", 1.1)
	mass31("bread", 82.3)

	mass4 := addInShopList()
	mass3("fish", 11.2)
	mass3("coffee", 2.1)
	mass3("a", 3.3)

	sumList := addListShopsList()
	sumList(&mass)
	sumList(&mass2)
	sumList(&mass3)
	sumList(&mass4)


	sumList1 := addListShopsList()
	sumList1(&mass1)
	sumList1(&mass21)
	sumList1(&mass31)

	fmt.Println(printBalance(clientList("Putin", sumList(nil))))
	fmt.Println(printBalance(clientList("Rich", sumList1(nil))))
	fmt.Println(printBalance(clientList("Mask", sumList1(nil))))
	fmt.Println(printBalance(clientList("Elan", sumList1(nil))))


	sortClientsName(clientList("", sumList1(nil)))
	sortClientsNameReverse(clientList("", sumList1(nil)))
	sortClientsBalanceReverse(clientList("", sumList1(nil)))
}
