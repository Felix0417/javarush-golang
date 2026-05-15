package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Print("0 -1")
		return
	}

	var max int
	fmt.Scan(&max)
	maxIdx := 0
	
	for i := 1; i < n; i++ {
		var x int
		fmt.Scan(&x)

		if x > max {
			max = x
			maxIdx = i
		}
	}

	fmt.Printf("%d %d\n", max, maxIdx)
}
