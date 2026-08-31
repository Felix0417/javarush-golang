package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s) // читаем строку как один токен (без пробелов)

	r := []rune(s) // работаем с Unicode-символами, а не с байтами
	fmt.Println(len(r))

	// TODO: Разверните слайс рун r "на месте" обменами элементов (не более len(r)/2 обменов).
	// TODO: После разворота выведите результат строго через string(r).
	for i := 0; i < len(r)/2; i++ {
		head := r[i]
		tailPos := len(r) - i - 1
		r[i] = r[tailPos]
		r[tailPos] = head
	}

	fmt.Println(string(r))
}
