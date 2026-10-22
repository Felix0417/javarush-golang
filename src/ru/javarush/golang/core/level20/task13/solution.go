package main

import "fmt"

func add(p *int, x int) {
	// TODO: Увеличьте значение по адресу p на x (через разыменование указателя).
	*p += x
}

func main() {
	var firstDeposit, secondDeposit int
	fmt.Scan(&firstDeposit, &secondDeposit)

	sumPtr := new(int) // копилка: указатель на сумму, стартует с zero value = 0

	add(sumPtr, firstDeposit)
	add(sumPtr, secondDeposit)

	fmt.Println(*sumPtr)
}
