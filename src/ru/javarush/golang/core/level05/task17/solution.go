package main

import "fmt"

func main() {
	var operationsCount int
	fmt.Scan(&operationsCount)

	totalSum := 0

	for i := 0; i < operationsCount; i++ {
		var op int
		fmt.Scan(&op)

		totalSum += op
	}

	fmt.Println(totalSum)
}
