package main

import (
	"fmt"
)

func main() {
	var income, expense, tax int
	fmt.Scan(&income, &expense, &tax)

	// Печатаем готовую строку отчёта из "мини-API" (package main).
	fmt.Printf("%s\n", PrintBalanceReport(income, expense, tax))
}
