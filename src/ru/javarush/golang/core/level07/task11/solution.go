package main

import "fmt"

func main() {
	var unitPriceCents, quantity, discountPercent int
	fmt.Scan(&unitPriceCents, &quantity, &discountPercent)

	subtotalCents := unitPriceCents * quantity

	discountCents := subtotalCents * discountPercent / 100

	totalCents := subtotalCents - discountCents

	eur := totalCents / 100
	cen := totalCents % 100

	fmt.Printf("%d %02d", eur, cen)
}
