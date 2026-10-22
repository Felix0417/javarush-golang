package main

import (
	"fmt"
	"strings"
)

func main() {
	var rawTitle string
	fmt.Scan(&rawTitle) // читаем одним токеном, без строки целиком

	// '_' считается пробелом
	s := strings.ReplaceAll(rawTitle, "_", " ")

	// Нормализация: первый шаг оставляем, остальное — доделать
	s = strings.TrimSpace(s)

	// TODO: Нормализуйте пробелы строго в последовательности: strings.Fields → strings.Join(parts, " ").
	// TODO: После полной нормализации, если итоговая строка пустая — выведите ровно EMPTY.
	s = strings.Join(strings.Fields(s), " ")

	if s == "" {
		fmt.Println("EMPTY")
		return
	}

	fmt.Println(s)
}
