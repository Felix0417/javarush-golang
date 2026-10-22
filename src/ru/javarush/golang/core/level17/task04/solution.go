package main

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrBadInput = errors.New("bad input")

var lastResult int // результат успешного run(), чтобы main мог вывести RESULT

func parseCommand(cmd string) error {
	// TODO: Реализуйте проверку команды: допускается только "div".
	// TODO: Для неизвестной команды верните ошибку с wrapping и первопричиной ErrBadInput.
	if cmd != "div" {
		return fmt.Errorf("command: %w", ErrBadInput)
	}
	return nil
}

func parseNumber(name, s string) (int, error) {
	// TODO: Реализуйте парсинг числа через strconv.Atoi.
	// TODO: Ошибку парсинга оберните с контекстом имени (name) через wrapping.
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse %s: %w", name, err)
	}
	return n, nil
}

func divide(a, b int) (int, error) {
	// TODO: Реализуйте целочисленное деление a/b.
	// TODO: При b == 0 верните ошибку с wrapping и первопричиной ErrBadInput.
	if b == 0 {
		return 0, fmt.Errorf("divide: %w", ErrBadInput)
	}
	return a / b, nil
}

func rootCause(err error) error {
	// TODO: Реализуйте снятие всех слоёв wrapping через errors.Unwrap в цикле (без рекурсии).
	// TODO: Верните самую внутреннюю (корневую) ошибку.
	for {
		e := errors.Unwrap(err)
		if e == nil {
			break
		}
		err = e
	}
	return err
}

func run() error {
	var cmd, aStr, bStr string
	fmt.Scan(&cmd, &aStr, &bStr)

	if err := parseCommand(cmd); err != nil {
		// TODO: Добавьте единый верхнеуровневый контекст run через wrapping.
		return fmt.Errorf("run: %w", err)
	}

	a, err := parseNumber("a", aStr)
	if err != nil {
		// TODO: Добавьте единый верхнеуровневый контекст run через wrapping.
		return fmt.Errorf("run: %w", err)
	}

	b, err := parseNumber("b", bStr)
	if err != nil {
		// TODO: Добавьте единый верхнеуровневый контекст run через wrapping.
		return fmt.Errorf("run: %w", err)
	}

	res, err := divide(a, b)
	if err != nil {
		// TODO: Добавьте единый верхнеуровневый контекст run через wrapping.
		return fmt.Errorf("run: %w", err)
	}

	lastResult = res
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		fmt.Printf("ROOT: %v\n", rootCause(err))
		return
	}

	fmt.Printf("RESULT: %d\n", lastResult)
}
