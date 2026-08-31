package main

import "fmt"

// OrderStatus — enum-тип статуса заказа (по требованию задачи).
type OrderStatus int

const (
	Created   OrderStatus = iota // 0
	Paid                         // 1
	Shipped                      // 2
	Delivered                    // 3
	Cancelled                    // 4
)

// statusName возвращает человекочитаемое имя статуса, неизвестное — "unknown".
func statusName(s OrderStatus) string {
	// TODO: Реализуйте отображение всех известных статусов в их строковые имена через switch.
	// TODO: Для неизвестных значений обязательно возвращайте "unknown" через ветку default.
	switch s {
	case Created:
		return "created"
	case Paid:
		return "paid"
	case Shipped:
		return "shipped"
	case Delivered:
		return "delivered"
	case Cancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

// isFinal — финальные статусы: Delivered и Cancelled, остальные (включая unknown) — не финальные.
func isFinal(s OrderStatus) bool {
	// TODO: Реализуйте проверку финальности статуса без "магических чисел":
	// TODO: сравнивайте только с enum-константами (Delivered/Cancelled), неизвестные считаются не финальными.
	if s == Delivered || s == Cancelled {
		return true
	}
	return false
}

func main() {
	var statusCode int
	fmt.Scan(&statusCode)

	status := OrderStatus(statusCode)

	fmt.Println(statusName(status))
	if isFinal(status) {
		fmt.Println("final")
	} else {
		fmt.Println("not_final")
	}
}
