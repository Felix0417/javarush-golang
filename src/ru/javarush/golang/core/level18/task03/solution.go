package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// TODO: Добавьте отдельный defer для перевода строки так, чтобы newline печатался после всех чисел.
	defer fmt.Println()

	for i := 0; i < n; i++ {
		// TODO: На каждой итерации цикла используйте defer fmt.Print(i), чтобы числа напечатались в обратном порядке (LIFO).
		defer fmt.Print(i)
	}
}
