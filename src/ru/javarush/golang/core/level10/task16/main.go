package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	// Собираем результат из заранее подготовленных частей.
	fmt.Print(left + s + right)
}
