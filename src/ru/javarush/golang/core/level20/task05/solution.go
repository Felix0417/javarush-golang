package main

import "fmt"

func main() {
	var message string
	fmt.Scan(&message) // читаем одним токеном (без пробелов)

	runeCount := 0
	// TODO: Посчитайте количество рун (Unicode-символов) в строке message через цикл for range и запишите результат в runeCount.
	for range message {
		runeCount++
	}
	fmt.Println(len(message)) // длина в байтах (UTF-8)
	fmt.Println(runeCount)    // количество рун
}
