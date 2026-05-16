package main

import "fmt"

func main() {
	total := 0
	fmt.Scan(&total)

	discount := 0
	if total >= 1000 {
		discount = 100
	} else {
		discount = 0
	}

	finalTotal := total - discount

	fmt.Printf("%d %d\n", discount, finalTotal)
}
