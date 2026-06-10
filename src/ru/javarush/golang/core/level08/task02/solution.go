package main

import "fmt"

const add = 5          // TODO: объявите untyped-константу add со значением 5 (без указания типа)
const addTyped int = 5 // TODO: объявите typed-константу addTyped типа int со значением 5

func main() {
	var x int64
	fmt.Scan(&x)

	// TODO: посчитайте сумму x + add без явного приведения add
	sum1 := x + add

	// TODO: посчитайте сумму x + int64(addTyped) с явным приведением typed-константы к int64
	sum2 := x + int64(addTyped)

	fmt.Printf("untyped=%d\n", sum1)
	fmt.Printf("typed=%d\n", sum2)
}
