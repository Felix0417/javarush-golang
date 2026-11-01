package main

import (
	"fmt"
	"strings"
)

// PromoCode — пользовательский тип на базе string.
type PromoCode string

// Normalized приводит код к "нормальному виду".
func (c PromoCode) Normalized() PromoCode {
	// TODO: Нормализуйте промокод: уберите пробелы по краям и приведите к верхнему регистру.
	trimmed := strings.ToUpper(strings.TrimSpace(string(c)))
	return PromoCode(trimmed)
}

// DiscountPercent определяет скидку по коду.
func (c PromoCode) DiscountPercent() int {
	// TODO: Определите процент скидки по нормализованному промокоду (10/20/0).
	promoCode := c.Normalized()
	if promoCode == "SALE10" {
		return 10
	} else if promoCode == "SALE20" {
		return 20
	}
	return 0
}

// Purchase — покупка с суммой и промокодом.
type Purchase struct {
	Amount int
	Code   PromoCode
}

// FinalAmount считает итоговую сумму по формуле из условия (целочисленно).
func (p Purchase) FinalAmount() int {
	// TODO: Рассчитайте итоговую сумму с учётом скидки, используя Purchase.Code.DiscountPercent().
	return p.Amount * (100 - p.Code.DiscountPercent()) / 100
}

func main() {
	var amount int
	var rawCode string

	// По условию читаем два значения: amount и rawCode.
	fmt.Scan(&amount, &rawCode)

	p := Purchase{
		Amount: amount,
		Code:   PromoCode(rawCode),
	}

	fmt.Println(p.FinalAmount())
}
