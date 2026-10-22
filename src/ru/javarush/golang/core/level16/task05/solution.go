package main

import (
	"errors"
	"fmt"
)

// Divide инкапсулирует целочисленное деление и возвращает ошибку при делении на ноль.
func Divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		// TODO: Верните 0 и ошибку, созданную через errors.New, с одним фиксированным текстом сообщения.
		return 0, errors.New("Divide by zero is not allowed!")
	}

	// TODO: Реализуйте успешное целочисленное деление dividend/divisor и верните nil в качестве ошибки.
	return dividend / divisor, nil
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	res, err := Divide(a, b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("result:", res)
}
