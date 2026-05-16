package main

import "fmt"

func main() {
	var x, y, z float64
	fmt.Scan(&x, &y, &z)

	minVal := x
	if y < minVal {
		minVal = y
	}
	if z < minVal {
		minVal = z
	}

	maxVal := x
	if y > maxVal {
		maxVal = y
	}
	if z > maxVal {
		maxVal = z
	}

	avg := (x + y + z) / 3

	stable := (maxVal - minVal) <= 1.0

	fmt.Printf("min=%v\n", minVal)
	fmt.Printf("max=%v\n", maxVal)
	fmt.Printf("avg=%v\n", avg)
	fmt.Printf("stable=%v\n", stable)
}
