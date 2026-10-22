package main

import (
	"fmt"
	"strconv"
)

// ParseAge парсит возраст из строки.
// Важно: внутри используем именно strconv.Atoi, а ошибку формируем через fmt.Errorf без %w.
func ParseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		// TODO: Добавьте контекст к ошибке парсинга возраста: исходную строку выведите через %q, а ошибку Atoi вставьте через %v (без %w).

		return 0, fmt.Errorf("ошибка парсинга возраста %q: %v", s, err)
	}

	// TODO: Проверьте, что age находится в диапазоне 0..130 (включительно); иначе верните 0 и ошибку через fmt.Errorf с указанием числового значения возраста.
	if age < 0 || age > 130 {
		return 0, fmt.Errorf("%d", age)
	}
	return age, nil
}

func main() {
	var s string
	fmt.Scan(&s)

	age, err := ParseAge(s)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("age: %d\n", age)
}
