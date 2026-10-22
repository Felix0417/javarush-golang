package main

import "fmt"

func normalizeOrder(a, b *int) bool {
	// TODO: Если один из указателей nil — ничего не менять и вернуть false.
	if a == nil || b == nil {
		return false
	}

	// TODO: Сравнить *a и *b; если *a > *b — поменять значения по адресам местами и вернуть true.
	if *a > *b {
		*a, *b = *b, *a
		return true
	}

	// TODO: Если *a <= *b — ничего не менять и вернуть false.
	return false
}

func main() {
	var leftValue, rightValue int
	fmt.Scan(&leftValue, &rightValue)

	swapped := normalizeOrder(&leftValue, &rightValue)

	fmt.Printf("min=%d max=%d\n", leftValue, rightValue)
	fmt.Printf("swapped=%v\n", swapped)
}
