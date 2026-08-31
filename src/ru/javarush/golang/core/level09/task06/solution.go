package main

import "fmt"

// minMax3 возвращает сначала минимум, затем максимум из трёх чисел.
func minMax3(a, b, c int) (int, int) {
	// TODO: Реализуйте поиск минимума и максимума среди a, b, c.
	// TODO: Используйте только сравнения (<, >) и ветвления if/else (без циклов, массивов и слайсов).
	// TODO: Важно: сначала верните min, затем max (порядок результатов важен).

	// Временная заглушка: возвращает значения, которые не являются корректными для общего случая.

	return calcMin(calcMin(a, b), c), calcMax(calcMax(a, b), c)
}
func calcMin(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

func calcMax(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func main() {
	var measureA, measureB, measureC int
	fmt.Scan(&measureA, &measureB, &measureC)

	mn, mx := minMax3(measureA, measureB, measureC)

	fmt.Printf("min=%d\n", mn)
	fmt.Printf("max=%d\n", mx)
}
