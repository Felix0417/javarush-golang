package main

import "fmt"

func main() {
	var (
		firstMeasure  float64
		secondMeasure float64
		thirdMeasure  float64
	)

	fmt.Scan(&firstMeasure, &secondMeasure, &thirdMeasure)

	var average float64 = (firstMeasure + secondMeasure + thirdMeasure) / 3

	fmt.Printf("avg=%.2f\n", average)
	fmt.Printf("avg=%.17f\n", average)
}
