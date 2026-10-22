package main

import (
	"errors"
	"fmt"
	"os"
)

// Sentinel-ошибки: их будем распознавать через errors.Is, а не по тексту.
var ErrUnknownOp = errors.New("unknown operation")
var ErrDivByZero = errors.New("division by zero")

func Calc(left, right int, op string) (int, error) {
	// TODO: Реализуйте поддерживаемые операции (+, -, *, /) и корректные ошибки.
	// TODO: Для ErrUnknownOp и ErrDivByZero возвращайте ошибки, обёрнутые через fmt.Errorf(... %w ...).
	switch op {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			// TODO: Верните ошибку деления на ноль, обёрнутую через %w.
			return 0, fmt.Errorf("calc: %w", ErrDivByZero)
		}
		return left / right, nil
	default:
		// TODO: Для неизвестной операции возвращайте ErrUnknownOp, обёрнутую через %w.
		return 0, fmt.Errorf("calc: %w", ErrUnknownOp)
	}
}

func main() {
	in := os.Stdin
	out := os.Stdout

	var left, right int
	var op string
	fmt.Fscan(in, &left, &op, &right)

	res, err := Calc(left, right, op)
	if err != nil {
		// TODO: Выведите нужные сообщения об ошибках, определяя причину через errors.Is.
		if errors.Is(err, ErrUnknownOp) {
			// TODO: Выведите "unsupported operation".
			fmt.Fprint(out, "unsupported operation")
		} else if errors.Is(err, ErrDivByZero) {
			// TODO: Выведите "division by zero".
			fmt.Fprint(out, "division by zero")
		}
		return
	}

	fmt.Fprint(out, res)
}
