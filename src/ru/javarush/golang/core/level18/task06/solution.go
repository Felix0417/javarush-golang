package main

import (
	"errors"
	"fmt"
	"sync"
)

var mu sync.Mutex
var balance int

func AddBalance(delta int) error {
	// TODO: Защитите работу с balance мьютексом mu: сделайте mu.Lock() и сразу следующей строкой defer mu.Unlock().
	mu.Lock()
	defer mu.Unlock()

	if delta == 0 {
		// TODO: Верните ошибку с точным текстом "delta is zero".
		return errors.New("delta is zero")
	}

	// TODO: Прибавьте delta к глобальному balance и верните nil.
	balance += delta
	return nil
}

func main() {
	balance = 10

	var delta int
	fmt.Scan(&delta)

	err1 := AddBalance(delta)
	err2 := AddBalance(1)

	mu.Lock()
	b := balance // читаем баланс под мьютексом
	mu.Unlock()

	fmt.Printf("err1: %v\n", err1)
	fmt.Printf("err2: %v\n", err2)
	fmt.Printf("balance: %d\n", b)
}
