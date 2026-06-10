package main

import "fmt"

func main() {
	var totalUnits, partsCount int
	fmt.Scan(&totalUnits, &partsCount)

	integerQuotient := totalUnits / partsCount

	floatQuotient := float64(totalUnits) / float64(partsCount)

	fmt.Println(integerQuotient)
	fmt.Println(floatQuotient)
}
