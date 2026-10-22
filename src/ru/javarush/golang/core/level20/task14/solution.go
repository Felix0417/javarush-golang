package main

import (
	"fmt"
	"strconv"
)

func finalPrice(price int, discount *int) int {
	// nil означает: скидки нет
	if discount == nil {
		return price
	}

	// TODO: Реализуйте расчёт итоговой цены со скидкой и защиту от отрицательного результата (итог не должен быть меньше 0).
	// TODO: Разыменовывайте discount (*discount) только после проверки discount на nil.

	total := price - *discount
	if total < 0 {
		return 0
	}

	return total
}

func main() {
	var price int
	var discountText string
	fmt.Scan(&price, &discountText)

	var discount *int
	if discountText != "nil" {
		d := 0

		// TODO: Распарсьте discountText в int через strconv.Atoi и используйте полученное значение как скидку.
		// TODO: Передайте в finalPrice указатель на распарсенное значение.
		d, _ = strconv.Atoi(discountText)

		discount = &d
	}

	fmt.Println(finalPrice(price, discount))
}
