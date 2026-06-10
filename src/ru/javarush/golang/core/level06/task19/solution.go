package main

import "fmt"

func main() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)

	for i := 0; i < n; i++ {
		// TODO: Обновите пару (x, y) по правилу из условия одной строкой (множественное присваивание), без временных переменных.
		x, y = y, x+y
	}

	fmt.Printf("x=%d y=%d\n", x, y)
}
