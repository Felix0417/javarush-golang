package main

import "fmt"

func main() {
	var found bool

	for i := 0; i < 5; i++ {
		var x int
		fmt.Scan(&x)

		if x == 10 {
			found = true
			break
		}
	}

	fmt.Println(found)
}
