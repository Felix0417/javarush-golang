package main

import "fmt"

func main() {
	i := 17

	i64 := int64(17)

	u := uint(17)

	fmt.Printf("%v (%T)\n", i, i)
	fmt.Printf("%v (%T)\n", i64, i64)
	fmt.Printf("%v (%T)\n", u, u)
}
