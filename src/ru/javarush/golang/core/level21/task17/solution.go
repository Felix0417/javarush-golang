package main

import (
	"errors"
	"fmt"
	"strings"
)

// User — модель пользователя с инвариантами на полях.
type User struct {
	Name string
	Age  int
}

func ValidateUser(u User) error {
	// TODO: Реализуйте полную валидацию пользователя:
	// TODO: 1) имя не пустое после strings.TrimSpace
	// TODO: 2) возраст в диапазоне от 1 до 150 включительно
	trimmedName := strings.TrimSpace(u.Name)

	if trimmedName == "" {
		return errors.New("name is empty")
	}

	if u.Age < 1 || u.Age > 150 {
		return errors.New("age must be between 1 and 150")
	}

	return nil
}

func main() {
	var nameInput string
	var ageInput int

	fmt.Scan(&nameInput, &ageInput)

	u := User{
		Name: nameInput,
		Age:  ageInput,
	}

	if err := ValidateUser(u); err != nil {
		fmt.Println("INVALID:", err.Error())
		return
	}

	fmt.Println("OK")
}
