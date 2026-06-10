package main

import (
	"fmt"
)

func main() {
	var incomeCents, expenseCents int64
	fmt.Scan(&incomeCents, &expenseCents)

	// Баланс тоже строго int64 (никаких int/float).
	var balanceCents int64
	balanceCents = incomeCents - expenseCents

	// Первая строка "квитанции": значение баланса + его тип через %T.
	fmt.Printf("balanceCents=%d type=%T\n", balanceCents, balanceCents)
	// TODO: выведите тип именно balanceCents через %T (не используйте литералы)

	// Вторая строка: статус по знаку баланса (строго одно слово).
	if balanceCents == 0 {
		fmt.Println("zero")
	} else if balanceCents > 0 {
		fmt.Println("profit")
	} else {
		fmt.Println("loss")
	}
	// TODO: определите статус по balanceCents через if/else if/else и выведите profit/loss/zero
}
