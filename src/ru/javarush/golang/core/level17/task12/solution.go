package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// sentinel: класс ситуации "возраст вне допустимых границ"
var ErrAgeOutOfRange = errors.New("age out of range")

func ParseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		// TODO: вернуть ошибку с контекстом и оборачиванием исходной причины через %w,
		// TODO: чтобы *strconv.NumError можно было извлечь через errors.As в main
		return 0, fmt.Errorf("parse age: %w", err)
	}

	if age < 0 || age > 130 {
		// TODO: вернуть ошибку с контекстом и оборачиванием ErrAgeOutOfRange через %w,
		// TODO: чтобы errors.Is распознавал ситуацию "вне диапазона"
		return 0, fmt.Errorf("age out of range: %w", ErrAgeOutOfRange)
	}

	return age, nil
}

func main() {
	var s string
	fmt.Fscan(os.Stdin, &s)

	age, err := ParseAge(s)
	if err == nil {
		fmt.Printf("AGE: %d\n", age)
		return
	}

	// TODO: реализовать требуемую логику вывода:
	// TODO: 1) BAD_INPUT, если errors.As распознаёт *strconv.NumError
	// TODO: 2) OUT_OF_RANGE, если errors.Is распознаёт ErrAgeOutOfRange
	// TODO: 3) иначе ERROR: <текст ошибки>
	// TODO: порядок проверок важен: сначала errors.As, затем errors.Is
	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		fmt.Println("BAD_INPUT")
		return
	}

	if errors.Is(err, ErrAgeOutOfRange) {
		fmt.Println("OUT_OF_RANGE")
		return
	}

	fmt.Printf("ERROR: %v\n", err)
}
