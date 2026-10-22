package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	p := new(int)

	var initial int
	// TODO: Зафиксируйте начальное значение *p в переменную initial (до записи x).
	initial = *p

	// TODO: Запишите считанное число x в *p.
	*p = x

	// TODO: Увеличьте *p на 1 через разыменование (не используйте ++).
	*p = *p + 1

	fmt.Println(initial, *p)
}
