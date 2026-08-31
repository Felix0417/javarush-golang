package main

import "fmt"

func main() {
	var inputWord string
	fmt.Scan(&inputWord) // читаем одно "слово" (токен)

	// Важно: нельзя индексировать пустую строку.
	if len(inputWord) == 0 {
		fmt.Println("EMPTY")
		return
	}

	n := len(inputWord) // длина именно в байтах
	fmt.Println(n)

	// TODO: Реализуйте вывод оставшихся 4 строк для непустой строки:
	// TODO: (1) числовое значение первого байта, (2) первый байт как символ,
	// TODO: (3) числовое значение последнего байта, (4) последний байт как символ.
	fmt.Printf("%d\n", inputWord[0])
	fmt.Printf("%c\n", inputWord[0])
	fmt.Printf("%d\n", inputWord[n-1])
	fmt.Printf("%c\n", inputWord[n-1])
}
