package main

import (
	"errors"
	"fmt"
)

// Sentinel-ошибка уровня пакета: распознаём класс ошибки через errors.Is.
var ErrInvalidLogin = errors.New("invalid login")

func ValidateLogin(login string) error {
	// TODO: Реализуйте валидацию логина: он не должен быть пустым и должен быть длиной минимум 3 символа.
	// TODO: При невалидном логине верните ошибку с контекстом, которая оборачивает ErrInvalidLogin через fmt.Errorf с %w.
	if login == "" || len(login) < 3 {
		return fmt.Errorf("validate login: %w", ErrInvalidLogin)
	}

	return nil
}

func main() {
	var login string
	fmt.Scan(&login)

	err := ValidateLogin(login)
	if err == nil {
		fmt.Println("OK")
		return
	}

	// TODO: Реализуйте распознавание ErrInvalidLogin через errors.Is(err, ErrInvalidLogin).
	// TODO: Если класс ошибки ErrInvalidLogin — выведите INVALID.
	// TODO: Если ошибка другая — выведите ERROR: <текст ошибки>.
	if errors.Is(err, ErrInvalidLogin) {
		fmt.Println("INVALID")
		return
	}
	fmt.Printf("ERROR: %v\n", err)
}
