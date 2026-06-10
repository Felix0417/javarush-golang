package main

import "fmt"

func main() {
	var from, to int
	fmt.Scan(&from, &to)

	if from > to {
		from, to = to, from
	}

	fmt.Printf("range: %d..%d\n", from, to)
}
