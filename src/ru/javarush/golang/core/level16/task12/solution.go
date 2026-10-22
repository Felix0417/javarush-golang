package main

import (
	"fmt"
	"strconv"
)

func run() error {
	var n int
	fmt.Scan(&n)

	// TODO: Проверьте, что N > 0. При некорректном N верните ошибку (error) с понятным текстом.
	if n <= 0 {
		return fmt.Errorf("incorrect value: %d", n)
	}

	var maxVal int

	for i := 1; i <= n; i++ {
		var tok string
		fmt.Scan(&tok)

		v, err := strconv.Atoi(tok)

		// TODO: Сразу после strconv.Atoi проверьте err и при ошибке немедленно верните её наверх.
		// TODO: В ошибке укажите номер элемента (i) и исходный токен (в кавычках через %q) в требуемом формате.
		// Сейчас ошибка парсинга игнорируется, а значение подменяется на 0.
		if err != nil {
			return fmt.Errorf("parse #%d %q: %v", i, tok, err)
		}
		if i == 1 {
			maxVal = v
		}

		// TODO: Реализуйте корректное вычисление максимума для всех целых чисел, включая отрицательные.
		if v > maxVal {
			maxVal = v
		}
	}

	fmt.Printf("max: %d\n", maxVal)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
