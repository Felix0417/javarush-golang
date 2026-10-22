package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var pairsCount int
	fmt.Fscan(in, &pairsCount)

	// По требованиям: создаём map через make до первой записи.
	prices := make(map[string]int)

	for i := 0; i < pairsCount; i++ {
		var productCode string
		var price int
		fmt.Fscan(in, &productCode, &price)
		prices[productCode] = price
	}

	var queryCode string
	fmt.Fscan(in, &queryCode)

	// TODO: Реализуйте корректную проверку наличия ключа через ok-идиому, чтобы не путать отсутствие товара и цену 0.
	// TODO: Выведите FOUND <value>, если товар есть в справочнике (даже если цена 0), иначе выведите MISSING.

	// Намеренно упрощённая (и неверная) проверка: цена 0 ошибочно считается "нет товара".
	price, ok := prices[queryCode]
	if ok {
		fmt.Printf("FOUND %d", price)
	} else {
		fmt.Print("MISSING")
	}
}
