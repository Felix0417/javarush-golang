package main

import (
	"errors"
	"fmt"
)

var ErrUnknownOp = errors.New("unknown op")

var ErrDivisionByZero = errors.New("division by zero")

func ParseOp(op string) (string, error) {
	// TODO: Провалидируйте op: разрешены только "add" и "div". Для остальных верните ErrUnknownOp (sentinel).
	switch op {
	case "add", "div":
		return op, nil
	default:
		return "", ErrUnknownOp
	}
}

func computeValidated(op string, a, b int) int {
	switch op {
	case "add":
		// TODO: Реализуйте операцию сложения.
		return a + b
	case "div":
		// TODO: Реализуйте операцию целочисленного деления.
		// TODO: Эта функция должна предполагать, что деление на ноль уже отфильтровано до её вызова.
		return a / b
	default:
		panic("unreachable: op was validated")
	}
}

func main() {
	var op string
	var a, b int
	fmt.Scan(&op, &a, &b)

	op, err := ParseOp(op)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	if op == "div" && b == 0 {
		// TODO: Если b == 0, обработайте это как пользовательскую ошибку до вызова computeValidated:
		// TODO: напечатайте "error: <ошибка>" и завершите программу.
		fmt.Println("error:", ErrDivisionByZero)
		return
	}

	fmt.Println(computeValidated(op, a, b))
}
