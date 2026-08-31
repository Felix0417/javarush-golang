package main

import "fmt"

// minMax возвращает минимум и максимум; результаты названы в сигнатуре.
func minMax(a, b int) (min int, max int) {
	// TODO: Реализуйте корректное определение минимума и максимума для a и b.
	// TODO: Убедитесь, что перед выходом из функции min и max присвоены явно (включая случай a == b).

	if a < b {
		min = a
		max = b
	} else if a == b {
		min = a
		max = a
	} else {
		min = b
		max = a
	}

	// Naked return по требованию
	return
}

func main() {
	var heightA, heightB int
	fmt.Scan(&heightA, &heightB)

	min, max := minMax(heightA, heightB)
	fmt.Printf("%d %d", min, max)
}
