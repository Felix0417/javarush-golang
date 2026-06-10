package main

import "fmt"

func main() {
	var baseCost, extraUnits, multiplier int64
	fmt.Scan(&baseCost, &extraUnits, &multiplier)

	var estimateA int64 = baseCost + extraUnits*multiplier
	var estimateB int64 = (baseCost + extraUnits) * multiplier

	// TODO: Вычислите две оценки стоимости по условию задачи и присвойте их переменным estimateA и estimateB.

	fmt.Print(estimateA, " ", estimateB)
}
