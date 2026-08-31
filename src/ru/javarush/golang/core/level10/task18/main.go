package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	// Вывод строго в заданном формате: две строки.
	fmt.Printf("sum=%d\n", Sum(a, b))
	fmt.Printf("product=%d\n", Product(a, b))
}
