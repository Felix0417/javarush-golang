package main

import "fmt"

func main() {
	var (
		city   string
		days   int
		perDay int
	)

	// Ввод: одно слово + два целых числа.
	fmt.Scan(&city, &days, &perDay)

	total := days * perDay

	fmt.Printf("Trip to %s: %d days, $%d/day, total=$%d\n", city, days, perDay, total)
}
