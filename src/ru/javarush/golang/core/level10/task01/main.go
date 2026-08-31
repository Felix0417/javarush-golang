package main

import "fmt"

func main() {
	// main — “дирижёр”: связывает ввод, вычисление и вывод
	a, b := readTwoInts()
	fmt.Printf("sum=%d\n", sum(a, b))
}
