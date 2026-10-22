package main

import (
	"errors"
	"fmt"
)

// Sentinel-ошибки: сравнимы по ссылке и удобны для errors.Is.
var (
	ErrUnknownCommand = errors.New("unknown command")
	ErrDivByZero      = errors.New("division by zero")
)

func Compute(cmd string, a, b int) (int, error) {
	switch cmd {
	case "add":
		return a + b, nil
	case "div":
		if b == 0 {
			return 0, ErrDivByZero
		}
		// TODO: Реализуйте корректное целочисленное деление a / b и верните результат с nil-ошибкой.
		return a / b, nil
	default:
		// TODO: Реализуйте корректную ошибку неизвестной команды (sentinel + wrapping) в требуемом формате.
		return 0, fmt.Errorf("%w: %q", ErrUnknownCommand, cmd)
	}
}

func main() {
	var cmd string
	var a, b int
	fmt.Scan(&cmd, &a, &b)

	res, err := Compute(cmd, a, b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(res)
}
