package main

import "fmt"

func main() {
	sumParts, expected := 0.1+0.2, 0.3

	fmt.Printf("%.17f\n", sumParts)
	fmt.Printf("%.17f\n", expected)

	eq := sumParts == expected
	fmt.Println(eq)

	diff := sumParts - expected
	fmt.Printf("%.17f\n", diff)
}
