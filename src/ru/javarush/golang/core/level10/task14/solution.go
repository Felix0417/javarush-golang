package main

import "fmt"

// Глобальный префикс, задаётся до входа в main().
var prefix string

func init() {
	// TODO: Инициализируйте глобальную переменную prefix в init() без чтения stdin.
	// TODO: В init() нельзя использовать условия (if/switch) и циклы (for/range).
	prefix = ">> "
}

func main() {
	var word string
	fmt.Scan(&word) // читаем ровно одно слово

	// TODO: Используйте подготовленный глобальный prefix (инициализированный до main()).
	fmt.Println(prefix + word)
}
