package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errDivisionByZero = errors.New("division by zero")

func parseInt(s string) (int, error) {
	// TODO: Реализуйте парсинг строки в int через strconv.Atoi и оформите ошибку строго по формату из условия (без %w).
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse int %q: %v", s, err)
	}
	return n, nil
}

func compute(a, b int, op string) (int, error) {
	// TODO: Реализуйте поддержку операций "+", "-", "*", "/" и корректные ошибки (деление на ноль и неизвестная операция).
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errDivisionByZero
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %q", op)
	}
}

func main() {
	var sA, op, sB string
	fmt.Scan(&sA, &op, &sB)

	a, err := parseInt(sA)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	b, err := parseInt(sB)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	res, err := compute(a, b, op)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("result: %d\n", res)
}
