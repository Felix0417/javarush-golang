package main

import (
	"errors"
	"fmt"
	"strconv"
)

// parseAge парсит возраст из строки через strconv.Atoi.
// Если парсинг не удался — добавляем контекст "parse age" и НЕ теряем причину через %w.
func parseAge(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		// TODO: Оберните ошибку через fmt.Errorf с контекстом "parse age" и %w, чтобы errors.Unwrap возвращал причину.
		return 0, fmt.Errorf("parse age: %w", err)
	}
	return age, nil
}

func main() {
	var ageStr string
	fmt.Scan(&ageStr)

	age, err := parseAge(ageStr)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		fmt.Printf("CAUSE: %v\n", errors.Unwrap(err))
		return
	}

	fmt.Printf("AGE: %d\n", age)
}
