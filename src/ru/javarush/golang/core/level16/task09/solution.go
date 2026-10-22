package main

import (
	"fmt"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		// Формат сообщения об ошибке строго по условию.
		fmt.Printf("error: %v\n", err)
	}
}

func run() error {
	var s string

	// Читаем один токен и при ошибке сразу выходим наружу.
	_, err := fmt.Scan(&s)
	if err != nil {
		return err
	}

	// TODO: Распарсьте строку s в целое число через strconv.Atoi и сразу обработайте ошибку (ранний return err).
	// TODO: Если ошибок нет — вычислите удвоенное число и выведите результат в stdout.
	// TODO: Убедитесь, что при любой ошибке (чтение/парсинг) результат НЕ печатается.
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	fmt.Println(2 * n)
	return nil
}
