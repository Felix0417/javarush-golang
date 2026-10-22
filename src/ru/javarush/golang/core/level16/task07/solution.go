package main

import (
	"errors"
	"fmt"
)

// FinalPrice считает цену со скидкой.
// Внутри ничего не печатаем — только возвращаем результат и/или ошибку.
func FinalPrice(price, discount int) (int, error) {
	if price <= 0 {
		// Фиксированный текст для всех случаев некорректной цены.
		return 0, errors.New("invalid price")
	}

	// TODO: Добавьте валидацию discount (диапазон допустимых значений) и возврат ошибки с диагностикой (в тексте должно быть значение discount).
	// TODO: Реализуйте корректный расчёт финальной цены по условию задачи (целочисленная арифметика).
	if discount < 0 || discount > 100 {
		return 0, fmt.Errorf("invalid discount: %d", discount)
	}

	return price * (100 - discount) / 100, nil
}

func main() {
	var price, discount int
	fmt.Scan(&price, &discount)

	final, err := FinalPrice(price, discount)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	fmt.Printf("final: %d\n", final)
}
