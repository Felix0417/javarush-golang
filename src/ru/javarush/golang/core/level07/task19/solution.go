package main

import (
	"fmt"
	"math"
)

func main() {
	var originalPrice float64
	var discountPercent int
	fmt.Scan(&originalPrice, &discountPercent)

	discount := originalPrice * float64(discountPercent) / 100
	// TODO: Посчитайте discount по формуле из условия задачи (процент от originalPrice).

	afterDiscount := originalPrice - discount
	// TODO: Посчитайте afterDiscount как цену после применения скидки.

	roundedPrice := math.Round(afterDiscount*100) / 100
	// TODO: Округлите afterDiscount до 2 знаков как новое числовое значение (не через форматирование вывода).

	fmt.Printf("%.17f\n", afterDiscount)
	fmt.Printf("%.17f\n", roundedPrice)
	fmt.Printf("%.2f\n", roundedPrice)
}
