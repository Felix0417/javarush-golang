package main

import "fmt"

// error — это интерфейс, достаточно метода Error() string
type badInputError struct{}

func (badInputError) Error() string {
	// TODO: Верните фиксированное сообщение, чтобы badInputError реализовывал интерфейс error.
	return "bad input"
}

func main() {
	var code string
	fmt.Scan(&code)

	// Важно: err именно типа error.
	var err error

	// TODO: Реализуйте логику присваивания err:
	// TODO: - если введено "ok", err должен остаться nil (не создавайте значение ошибки)
	// TODO: - иначе err должен содержать значение вашего типа ошибки (badInputError)
	if code != "ok" {
		err = badInputError{}
	}

	// 1) Результат сравнения с nil (строго через err == nil).
	fmt.Println(err == nil)

	// 2) Как fmt печатает err (печать err напрямую).
	fmt.Println(err)

	// 3) Фактический (конкретный) тип значения внутри интерфейса error.
	fmt.Printf("%T\n", err)
}
