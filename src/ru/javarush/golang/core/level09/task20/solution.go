package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Sentinel-ошибки: текст должен совпадать с требованиями.
var (
	errDivisionByZero = errors.New("error: division by zero")
	errUnknownOp      = errors.New("error: unknown operator")
	errUnknownRule    = errors.New("error: unknown rule")
)

func main() {
	if err := run(); err != nil {
		fmt.Print(err.Error())
	}
}

func run() error {
	var aS, opS, bS, ruleCode string
	_, err := fmt.Scan(&aS, &opS, &bS, &ruleCode)
	if err != nil {
		return err
	}

	a, err := parseInt(aS)
	if err != nil {
		return err
	}

	b, err := parseInt(bS)
	if err != nil {
		return err
	}

	op, err := chooseOp(opS)
	if err != nil {
		return err
	}

	res, err := op(a, b)
	if err != nil {
		return err
	}

	rule, err := chooseRule(ruleCode)
	if err != nil {
		return err
	}

	res = apply(res, rule)
	fmt.Printf("result: %d", res)
	return nil
}

func parseInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func chooseOp(op string) (func(int, int) (int, error), error) {
	switch op {
	case "+":
		return func(a, b int) (int, error) { return a + b, nil }, nil
	case "-":
		// TODO: Реализуйте операцию вычитания и верните функцию, которая вычисляет a - b.
		return func(a, b int) (int, error) {
			return a - b, nil
		}, nil
	case "*":
		// TODO: Реализуйте операцию умножения и верните функцию, которая вычисляет a * b.
		return func(a int, b int) (int, error) {
			return a * b, nil
		}, nil
	case "/":
		// TODO: Реализуйте операцию деления и обработку деления на ноль (вернуть errDivisionByZero).
		return func(a int, b int) (int, error) {
			if b == 0 {
				return 0, errDivisionByZero
			}
			return a / b, nil
		}, nil
	default:
		return nil, errUnknownOp
	}
}

func chooseRule(code string) (func(int) int, error) {
	switch code {
	case "d":
		// TODO: Реализуйте правило "d" и верните функцию пост-обработки результата.
		return func(x int) int {
			return x * 2
		}, nil
	case "s":
		// TODO: Реализуйте правило "s" и верните функцию пост-обработки результата.
		return func(x int) int {
			return x * x
		}, nil
	case "n":
		return func(x int) int { return x }, nil
	default:
		return nil, errUnknownRule
	}
}

func apply(x int, r func(int) int) int {
	return r(x)
}
