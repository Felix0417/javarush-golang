package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	evenCount := 0 // счётчик чётных, по условию инициализируем до цикла

	for i := 0; i < n; i++ { // ровно n итераций
		var x int
		fmt.Scan(&x)

		if x%2 == 0 {
			evenCount++
		}
	}

	fmt.Println(evenCount)
}
