package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	sum := a + b
	diff := a - b
	prod := a * b

	fmt.Printf("sum=%d\n", sum)
	fmt.Printf("diff=%d\n", diff)
	fmt.Printf("prod=%d\n", prod)

	if b == 0 {
		fmt.Print("div=undefined\n")
		fmt.Print("mod=undefined\n")
	} else if b != 0 {
		fmt.Printf("div=%d\n", a/b)
		fmt.Printf("mod=%d\n", a%b)
	}

}
