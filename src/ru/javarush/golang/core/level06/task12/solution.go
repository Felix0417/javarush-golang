package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var sum float64   // сумма положительных
	var count float64 // количество положительных (тоже float64 по условию)
	var hasAny bool   // был ли хотя бы один положительный

	for i := 0; i < n; i++ {
		var x float64
		fmt.Scan(&x)

		if x > 0 {
			sum += x
			count = count + 1.0
			hasAny = true
		}

	}

	if !hasAny || count == 0 {
		fmt.Println(0)
		return
	}

	fmt.Println(sum / count)
}
