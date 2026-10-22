package main

import (
	"errors"
	"fmt"
)

// Sentinel-ошибка для распознавания причины через errors.Is
var ErrNegative = errors.New("negative balance")

func validateNonNegative(balance int) error {
	// TODO: Реализуйте проверку на отрицательный баланс: при balance < 0 верните ErrNegative, иначе nil.
	if balance < 0 {
		return ErrNegative
	}
	return nil
}

func main() {
	var balance int
	fmt.Scan(&balance)

	err := validateNonNegative(balance)
	if err != nil {
		// TODO: Оберните ошибку через %w (а не через %v), чтобы причина сохранялась в цепочке ошибок.
		err = fmt.Errorf("balance validation: %w", err)
	}

	// TODO: Определяйте причину через errors.Is(err, ErrNegative), а не сравнением err == ErrNegative и не по тексту ошибки.
	if errors.Is(err, ErrNegative) {
		fmt.Println("NEGATIVE")
	} else {
		fmt.Println("OK")
	}
}
