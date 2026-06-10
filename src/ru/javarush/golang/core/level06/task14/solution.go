package main

import "fmt"

func main() {
	var x int
	var s string

	fmt.Scan(&x, &s)

	fmt.Printf("x=%v type=%T\n", x, x)

	fmt.Printf("s=%v type=%T\n", s, s)
}
