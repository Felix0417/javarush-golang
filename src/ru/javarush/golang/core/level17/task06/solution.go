package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parseCount(s string) (int, error) {
	// Парсим строго через strconv.Atoi, как требуется условием.
	n, err := strconv.Atoi(s)
	if err != nil {
		// TODO: Оберните ошибку через fmt.Errorf с %w, чтобы сохранить исходную причину в цепочке.
		return 0, fmt.Errorf("parse count %q: %w", s, err)
	}
	return n, nil
}

func main() {
	var s string
	fmt.Scan(&s)

	n, err := parseCount(s)
	if err == nil {
		fmt.Println(n)
		return
	}

	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		// TODO: Если удалось извлечь *strconv.NumError, выведите "BAD_NUMBER: <значение>" (значение — numErr.Num) и завершитесь.
		fmt.Printf("BAD_NUMBER: %s\n", numErr.Num)
		return
	}

	// TODO: Обработайте неожиданные ошибки без panic. Формат вывода: "BAD_NUMBER: <значение>".
	fmt.Printf("BAD_NUMBER: %s\n", numErr.Num)
}
