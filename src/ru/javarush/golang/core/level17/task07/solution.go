package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field string
	Value string
}

func (e *ValidationError) Error() string {
	// TODO: сделайте человекочитаемое сообщение об ошибке, используя e.Field и e.Value.
	return fmt.Sprintf("%s - %s", e.Field, e.Value)
}

func validateUsername(name string) error {
	// TODO: реализуйте валидацию имени:
	// - если len(name) < 3, верните *ValidationError с нужными Field/Value
	// - иначе верните nil
	if len(name) < 3 {
		return &ValidationError{"username", name}
	}
	return nil
}

func main() {
	var name string
	fmt.Scan(&name)

	err := validateUsername(name)
	if err != nil {
		// TODO: оберните ошибку контекстом через fmt.Errorf("create user: %w", err)
		err = fmt.Errorf("create user: %w", err)
		// Пытаемся достать typed error (даже если ошибка будет обёрнута).
		var vErr *ValidationError
		if errors.As(err, &vErr) {
			// TODO: выведите строго: INVALID <Field> <Value>
			fmt.Printf("INVALID %s %s\n", vErr.Field, vErr.Value)
			return
		}

		// TODO: при необходимости обработайте ситуацию, когда это не ValidationError.
		fmt.Printf("INVALID %s\n", err)
		return
	}

	fmt.Println("ACCEPTED")
}
