package main

import "fmt"

func main() {
	var (
		item  string
		qty   int
		price int
	)

	fmt.Scan(&item, &qty, &price)

	fmt.Printf("Item=%s Qty=%d Price=%d\n", item, qty, price)

	var sum int
	sum = qty * price

	fmt.Printf("Sum=%d\n", sum)

	fmt.Printf("DEBUG: itemType=%T qtyType=%T priceType=%T sumType=%T\n", item, qty, price, sum)
}
