package main

import "fmt"

func main() {
	var usedSeconds, pricePerMinuteCents, baseFeeCents int
	fmt.Scan(&usedSeconds, &pricePerMinuteCents, &baseFeeCents)

	billableMinutes := usedSeconds / 60
	leftoverSeconds := usedSeconds % 60

	variableCents := billableMinutes * pricePerMinuteCents
	totalCents := baseFeeCents + variableCents

	euros := totalCents / 100
	cents := totalCents % 100

	fmt.Printf("%d %02d %d", euros, cents, leftoverSeconds)
}
