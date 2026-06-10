package main

import "fmt"

func main() {
	// TODO: Вычислите значение ratio прямо в коде как 2.0/3.0 (stdin читать нельзя).
	// TODO: Убедитесь, что переменная ratio имеет тип float64.
	var ratio float64 = 2.0 / 3.0

	// TODO: Выведите ratio с 2 знаками после точки, используя fmt.Printf и нужную форматную строку.
	fmt.Printf("Pretty: %.2f\n", ratio)

	// TODO: Выведите ratio с 17 знаками после точки, используя fmt.Printf и нужную форматную строку.
	fmt.Printf("Raw: %.17f\n", ratio)

	// TODO: Выведите то же значение переменной ratio (не изменяя её ради округления) с 17 знаками после точки.
	fmt.Printf("Still x: %.17f\n", ratio)
}
