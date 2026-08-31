package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parsePositiveInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	if n <= 0 {
		// TODO: Верните ошибку с корректным текстом (через errors.New), как требуется в условии.
		return 0, errors.New("number must be positive")
	}

	return n, nil
}

func main() {
	var rawCount string
	fmt.Scan(&rawCount)

	n, err := parsePositiveInt(rawCount)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		// TODO: При ошибке нужно сразу остановить выполнение (ранний выход), не печатая ничего больше.
		return
	}

	fmt.Printf("n = %d\n", n)
}
