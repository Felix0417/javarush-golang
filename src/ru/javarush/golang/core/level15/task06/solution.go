package main

import "fmt"

func main() {
	// false тут значит "заказ есть, но ещё не доставлен", поэтому при чтении нужен ok (v, ok := m[k])
	orders := map[string]bool{
		"ORD-100": true,
		"ORD-200": false,
	}

	var orderID string
	fmt.Scan(&orderID)

	// Требование: именно if с коротким объявлением + ok-идиома.
	if v, ok := orders[orderID]; ok {
		// TODO: Реализуйте вывод статуса для известного заказа.
		// TODO: Важно различать "ключа нет" и "значение false": unknown order зависит только от ok.
		// TODO: Должна быть выведена ровно одна строка без лишних переводов строк.
		if v {
			fmt.Print("delivered")
		} else {
			fmt.Print("not delivered")
		}
	} else {
		fmt.Print("unknown order")
	}
}
