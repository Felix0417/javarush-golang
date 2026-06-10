package main

import "fmt"

func main() {
	var itemsToPack, boxCapacity int
	fmt.Scan(&itemsToPack, &boxCapacity)

	var fullBoxes int = itemsToPack / boxCapacity
	var leftoverItems int = itemsToPack % boxCapacity

	// TODO: Посчитайте количество полностью заполненных коробок (целочисленное деление itemsToPack на boxCapacity).
	// TODO: Посчитайте количество оставшихся предметов (остаток от деления itemsToPack на boxCapacity).

	fmt.Print(fullBoxes, " ", leftoverItems)
}
