package main

import (
	"fmt"
	"os"
)

func validateSeats(seats int) error {
	// TODO: Реализуйте валидацию seats: допустимы только значения 1..300 (включительно).
	// TODO: Если значение некорректно — верните error с текстом в формате из условия (Go-стиль, без точки в конце, с подстановкой числа через %d).
	// TODO: Если значение корректно — верните nil.
	if seats <= 0 || seats > 300 {
		return fmt.Errorf("seats must be between 1 and 300, got %d", seats)
	}
	return nil
}

func main() {
	var seats int
	fmt.Fscan(os.Stdin, &seats)

	if err := validateSeats(seats); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ok")
}
