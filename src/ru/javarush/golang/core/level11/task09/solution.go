package main

import "fmt"

func main() {
	// a — nil-слайс (объявлен через var, памяти под элементы нет)
	var a []int

	// TODO: Создайте переменную b типа []int как пустой (empty) слайс через литерал, но НЕ nil.
	b := []int{}

	// TODO: Создайте переменную c типа []int как пустой (empty) слайс, но НЕ nil, с cap равным 3.
	c := make([]int, 0, 3)

	fmt.Printf("a: nil=%t len=%d cap=%d\n", a == nil, len(a), cap(a))
	fmt.Printf("b: nil=%t len=%d cap=%d\n", b == nil, len(b), cap(b))
	fmt.Printf("c: nil=%t len=%d cap=%d\n", c == nil, len(c), cap(c))
}
