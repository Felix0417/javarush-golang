package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Sentinel-ошибки = "классы" ошибок, которые распознаются через errors.Is.
var (
	ErrUnknownOp = errors.New("unknown op")
	ErrDivByZero = errors.New("division by zero")
)

func calc(op, aStr, bStr string) (int, error) {
	// TODO: Реализуйте проверку операции так, чтобы неизвестная операция возвращала ErrUnknownOp
	// TODO: Важно: неизвестную операцию нужно обрабатывать сразу, без парсинга чисел
	validateOp := op == "add" || op == "sub" || op == "mul" || op == "div"

	if !validateOp {
		return 0, ErrUnknownOp
	}

	a, err := strconv.Atoi(aStr)
	if err != nil {
		// TODO: Оберните ошибку парсинга через fmt.Errorf(... %w ...), чтобы *strconv.NumError оставался доступен через errors.As
		return 0, fmt.Errorf("parse %q: %w", aStr, err)
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		// TODO: Оберните ошибку парсинга через fmt.Errorf(... %w ...)
		return 0, fmt.Errorf("parse %q: %w", bStr, err)
	}

	// TODO: Реализуйте вычисления для add/sub/mul/div
	// TODO: Для div при b == 0 верните ErrDivByZero (класс ошибки через sentinel)
	// Сейчас намеренно неполная логика: возвращаем сложение для любой операции.

	switch op {
	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		return a * b, nil
	case "div":
		{
			if b == 0 {
				return 0, ErrDivByZero
			}
			return a / b, nil
		}
	default:
		return 0, ErrUnknownOp
	}
}

func main() {
	var op, aStr, bStr string
	fmt.Fscan(os.Stdin, &op, &aStr, &bStr)

	res, err := calc(op, aStr, bStr)
	if err == nil {
		fmt.Println(res)
		return
	}

	// TODO: На границе приложения нужно оборачивать ошибку через fmt.Errorf("calc: %w", err)
	err = fmt.Errorf("calc: %w", err)

	// Частично реализовано: класс неизвестной операции через errors.Is.
	// TODO: Дополните обработку ErrDivByZero через errors.Is
	// TODO: Дополните обработку BAD_NUMBER через errors.As и *strconv.NumError
	if errors.Is(err, ErrUnknownOp) {
		fmt.Println("UNKNOWN_OP")
		return
	}
	if errors.Is(err, ErrDivByZero) {
		fmt.Println("DIV_BY_ZERO")
		return
	}
	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		fmt.Printf("BAD_NUMBER: %v\n", numErr.Num)
		return
	}

	fmt.Printf("ERROR: %v\n", err)
}
