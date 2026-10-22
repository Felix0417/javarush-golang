package main

import "fmt"

func incLocal(x int) {
	// TODO: Увеличьте локальную копию аргумента x на 1 и выведите строку в формате: in func: %d
	x++
	fmt.Printf("in func: %d\n", x)
}

func main() {
	var startNumber int
	fmt.Scan(&startNumber)

	incLocal(startNumber)

	// TODO: Выведите значение startNumber из main в формате: in main: %d
	fmt.Printf("in main: %d\n", startNumber)
}
