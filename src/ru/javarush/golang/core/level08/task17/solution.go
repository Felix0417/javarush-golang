package main

import "fmt"

// Light — наш enum-тип для кода светофора.
type Light int

const (
	LightRed Light = iota
	LightYellow
	LightGreen
)

// lightText превращает код в человекочитаемый текст.
func lightText(l Light) string {
	// TODO: Реализуйте расшифровку кода светофора в текст.
	// TODO: Функция должна возвращать ровно одну из строк: "red", "yellow", "green" или "unknown".
	// TODO: В логике нельзя использовать числовые литералы 0/1/2 — сравнивайте только с LightRed/LightYellow/LightGreen.
	switch l {
	case LightRed:
		return "red"
	case LightYellow:
		return "yellow"
	case LightGreen:
		return "green"
	}
	return "unknown"
}

func main() {
	var lightCode int
	fmt.Scan(&lightCode)

	// По условию: преобразование делаем в main.
	fmt.Println(lightText(Light(lightCode)))
}
