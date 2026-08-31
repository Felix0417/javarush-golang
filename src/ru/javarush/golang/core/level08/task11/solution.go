package main

import "fmt"

type OrderStatus int

const (
	// TODO: Объявите константы OrderStatus через iota так, чтобы unknown соответствовал коду 0,
	// TODO: а "живые" статусы (created, paid, shipped) начинались с 1 (без ручной нумерации).
	unknown OrderStatus = iota
	created
	paid
	shipped
)

func main() {
	var code int
	fmt.Scan(&code)

	status := OrderStatus(code)

	// TODO: Выберите человекочитаемый текст через switch по status.
	// TODO: Нельзя сравнивать code/status с числовыми литералами 1/2/3 — используйте только enum-константы.
	// TODO: Для всех непонятных кодов (включая 0 и любые другие) выводите "unknown" через ветку default.
	switch status {
	case created:
		fmt.Println("created")
	case paid:
		fmt.Println("paid")
	case shipped:
		fmt.Println("shipped")
	default:
		fmt.Println("unknown")
	}
}
