package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	// TODO: Реализуйте целочисленное деление a / b с контрактом (int, error).
	// TODO: Если b == 0, верните (0, error) с текстом ошибки "division by zero" (без panic/recover).
	// TODO: Иначе верните результат деления и nil.
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func run() error {
	var a, b int
	fmt.Scan(&a, &b)

	res, err := divide(a, b)

	// TODO: Если err != nil — верните err наверх без подмены/скрытия.
	// TODO: Если ошибки нет — выведите строку "result: <число>" и переведите строку, затем верните nil.
	if err != nil {
		return err
	}
	fmt.Println("result:", res)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
