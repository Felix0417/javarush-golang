package main

import (
	"fmt"
	"math"
)

func main() {
	var totalWeight, maxBoxWeight, boxPrice float64
	fmt.Scan(&totalWeight, &maxBoxWeight, &boxPrice)

	rawBoxes := totalWeight / maxBoxWeight

	// TODO: Выведите PrettyBoxes: распечатайте rawBoxes через формат "%.0f" (без math.Ceil и без перевода в целый тип)
	fmt.Printf("%.0f\n", rawBoxes)

	boxesByRule := math.Ceil(rawBoxes)
	// TODO: Исправьте расчёт boxesByRule: по правилу нужно округлять вверх (math.Ceil), а затем печатать через "%.0f"
	fmt.Printf("%.0f\n", boxesByRule)

	total := boxesByRule * boxPrice

	// TODO: Выведите Total с двумя знаками после точки через формат "%.2f"
	fmt.Printf("%.2f\n", total)
}
