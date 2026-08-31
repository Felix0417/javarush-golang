package main

import "fmt"

// PrintBalanceReport — "публичная" функция мини-API: собирает итоговую строку отчёта.
func PrintBalanceReport(income, expense, tax int) string {
	balance := calcBalance(income, expense, tax) // формула живёт не в main.go
	return formatReport(income, expense, tax, balance)
}

// calcBalance — помощник для вычисления баланса (деталь реализации, не экспортируем).
func calcBalance(income, expense, tax int) int {
	// TODO: Реализуйте корректное вычисление баланса по условию задачи (не делайте это в main.go).
	return income - expense - tax
}

// formatReport — помощник для строгого форматирования отчёта через fmt.Sprintf.
func formatReport(income, expense, tax, balance int) string {
	// TODO: Сформируйте строку отчёта в строгом формате (ровно 4 поля) через fmt.Sprintf.
	return fmt.Sprintf("income %d expense %d tax %d balance %d", income, expense, tax, balance)
}
