package main

import "fmt"

func makeIntPtr(v int) *int {
	// TODO: Создайте локальную переменную-копию значения v и верните её адрес.
	// TODO: Убедитесь, что по возвращённому указателю хранится именно v.
	tmp := v
	return &tmp
}

func main() {
	var x int
	fmt.Scan(&x)

	p := makeIntPtr(x)

	x += 1000

	fmt.Print(x, " ", *p)
}
