package main

import (
	"errors"
	"fmt"
)

// Rule — правило преобразования силы.
type Rule func(int) int

func chooseRule(code string) (Rule, error) {
	switch code {
	case "n":
		return func(x int) int { return x }, nil
	// TODO: добавьте обработку кода "d" (удвоить) через анонимную функцию func(int) int
	// TODO: добавьте обработку кода "s" (возвести в квадрат) через анонимную функцию func(int) int
	case "d":
		return func(x int) int { return x * 2 }, nil
	case "s":
		return func(x int) int { return x * x }, nil
	default:
		// TODO: для неизвестного кода нужно вернуть nil и ошибку
		return nil, errors.New("unknown rule")
	}
}

func apply(x int, r Rule) int {
	return r(x)
}

func main() {
	var powerValue int
	var ruleCode string
	fmt.Scan(&powerValue, &ruleCode)

	r, err := chooseRule(ruleCode)
	if err != nil {
		fmt.Print("error: unknown rule")
		return
	}

	result := apply(powerValue, r)
	fmt.Printf("result: %d", result)
}
