package main

import "fmt"

func main() {
	var s string
	var i int
	fmt.Scan(&s, &i)

	// TODO: Реализуйте проверку, что i находится в диапазоне [0, len(s)) ДО обращения к s[i], чтобы избежать panic.
	// TODO: Проверка границ должна использовать сравнения только с 0 и len(s) (i < 0 или i >= len(s)).
	// TODO: Если индекс некорректный — выведите ровно OUT_OF_RANGE без лишних символов и завершите программу.
	// TODO: Если индекс корректный — получите байт строго через s[i] и выведите ровно: byte=<N> char=<C> (для символа используйте %c).

	if i < 0 || i >= len(s) {
		fmt.Print("OUT_OF_RANGE")
		return
	}

	symbol := s[i]

	fmt.Printf("byte=%d char=%c\n", symbol, symbol)

	// Временное поведение шаблона: всегда считаем индекс некорректным.

}
