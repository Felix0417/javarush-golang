package main

import (
	"errors"
	"fmt"
)

// ErrDivByZero — sentinel-ошибка для случая деления на ноль.
var ErrDivByZero = errors.New("division by zero")

// SafeDiv выполняет целочисленное деление a на b и возвращает ошибку при проблемах.
func SafeDiv(a, b int) (int, error) {
	// TODO: Реализуйте безопасное целочисленное деление a/b.
	// TODO: При b == 0 нужно вернуть (0, ErrDivByZero) и не использовать panic.
	// TODO: При корректном делении нужно вернуть (a / b, nil).
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	result, err := SafeDiv(a, b)

	// TODO: Реализуйте корректный вывод результата.
	// TODO: Если err != nil — вывести ровно одну строку через fmt.Println("error:", err).
	// TODO: Если err == nil — вывести ровно одну строку с result через fmt.Println(result).
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(result)
}
