package main

import (
	"errors"
	"fmt"
)

var errDivisionByZero = errors.New("division by zero")

func SafeDiv(total, divider int) (int, error) {
	// TODO: Реализуйте безопасное целочисленное деление total / divider.
	// TODO: Если divider == 0 — сделайте ранний возврат: (0, errDivisionByZero). panic использовать нельзя.
	if divider == 0 {
		return 0, errDivisionByZero
	}

	// TODO: Верните результат целочисленного деления и nil-ошибку.
	return total / divider, nil
}

func main() {
	var total, divider int
	fmt.Scan(&total, &divider)

	result, err := SafeDiv(total, divider)
	if err != nil {
		// TODO: При ошибке выведите ровно: error: division by zero
		fmt.Println("error:", err)
		return
	}

	fmt.Println(result)
}
