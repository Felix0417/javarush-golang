package main

import (
	"fmt"
	"os"
)

func deposit(balancePtr *int, amount int) bool {
	// Важно: при nil нельзя разыменовывать указатель.
	if balancePtr == nil {
		return false
	}

	// TODO: Реализуйте пополнение баланса по указателю (увеличьте значение на amount) и верните true, если изменение произошло.
	// TODO: Сейчас функция намеренно не меняет баланс и всегда сообщает, что изменений не было.
	*balancePtr += amount
	return true
}

func main() {
	var startBalance, topUpAmount int
	fmt.Fscan(os.Stdin, &startBalance, &topUpAmount)

	balanceValue := startBalance

	var ptr *int
	if startBalance >= 0 {
		ptr = &balanceValue
	} else {
		ptr = nil
	}

	changed := deposit(ptr, topUpAmount)

	fmt.Println(balanceValue)
	fmt.Printf("changed=%v\n", changed)
}
