package main

import "fmt"

func main() {
	// Входные данные: сумма счёта и процент чаевых (типы фиксированы требованием)
	var billAmount float64
	var tipPercent int
	fmt.Scan(&billAmount, &tipPercent)

	// TODO: Рассчитайте сумму чаевых (tip) на основе суммы счёта и процента чаевых.
	tip := billAmount * float64(tipPercent) / 100

	// TODO: Рассчитайте итоговую сумму (total) с учётом чаевых.
	total := billAmount + tip

	// Денежный вывод: всегда ровно 2 знака после точки
	fmt.Printf("Tip: %.2f\n", tip)
	fmt.Printf("Total: %.2f\n", total)
}
