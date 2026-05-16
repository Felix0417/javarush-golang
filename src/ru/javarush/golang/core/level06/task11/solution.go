package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println(0)
		return
	}

	var max int
	var hasMax bool

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)

		if !hasMax {
			max = x
			hasMax = true
		}

		if x > max {
			max = x
		}
	}

	fmt.Println(max)
}
