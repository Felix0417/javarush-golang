package main

import "fmt"

func main() {
	var steps int
	fmt.Scan(&steps)

	var sum64 float64 = 0.0
	var sum32 float32 = 0

	for i := 0; i < steps; i++ {
		sum64 += 0.1
		sum32 += 0.1
	}

	expected := float64(steps) / 10

	var delta64 float64
	if sum64 > expected {
		delta64 = sum64 - expected
	} else {
		delta64 = expected - sum64
	}

	var delta32 float64
	sum32InFloat64 := float64(sum32)
	if sum32InFloat64 > expected {
		delta32 = sum32InFloat64 - expected
	} else {
		delta32 = expected - sum32InFloat64
	}

	fmt.Printf("sum64=%.17f\n", sum64)
	fmt.Printf("sum32=%.17f\n", sum32)
	fmt.Printf("delta64=%.17f\n", delta64)
	fmt.Printf("delta32=%.17f\n", delta32)
}
