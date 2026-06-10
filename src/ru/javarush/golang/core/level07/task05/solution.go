package main

import "fmt"

func main() {
	defaultReading := 0.1
	var liteReading float32 = 0.1

	fmt.Printf("%f %T\n", defaultReading, defaultReading)
	fmt.Printf("%f %T\n", liteReading, liteReading)
}
