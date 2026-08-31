package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	// TODO: Получите первый байт строки выражением s[0] и сохраните в переменную b (тип byte).
	b := s[0]

	// TODO: Получите первый Unicode-символ (rune) через самую первую итерацию for range по строке.
	// TODO: После получения rune и его кода немедленно остановите цикл (например, через break).
	var r rune

	for _, v := range s {
		// TODO: Реализуйте извлечение первого rune из строки через range и остановку цикла.
		r = v
		break
	}

	fmt.Printf("byte=%d rune=%c code=%d\n", b, r, r)
}
