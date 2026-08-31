package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	doubled := 0
	// TODO: Вычислите удвоенное значение числа n и запишите его в doubled.

	doubled = 2 * n

	// Важно: не называем строковую переменную "fmt", иначе затенится пакет fmt.
	outFormat := "%d %d\n"
	fmt.Printf(outFormat, n, doubled)
}
