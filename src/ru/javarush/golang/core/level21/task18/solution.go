package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Coupon struct {
	Code    string
	Percent int
}

func ValidateCoupon(c Coupon) error {
	// TODO: Реализуйте полную валидацию инвариантов Coupon и возвращайте ПЕРВУЮ найденную ошибку.
	// TODO: Сообщения ошибок должны быть строго: "code is empty", "code is too short", "percent out of range".

	code := strings.TrimSpace(c.Code)
	if code == "" {
		return errors.New("code is empty")
	}

	// TODO: Добавьте проверку минимальной длины кода (>= 5).
	// TODO: Добавьте проверку процента скидки (1..90 включительно).
	if len([]rune(code)) < 5 {
		return errors.New("code is too short")
	}

	if c.Percent < 1 || c.Percent > 90 {
		return errors.New("percent out of range")
	}

	return nil
}

func NewCoupon(code string, percent int) (Coupon, error) {
	normalizedCode := strings.TrimSpace(code)

	c := Coupon{
		Code:    normalizedCode,
		Percent: percent,
	}

	// TODO: Убедитесь, что нормализация code происходит ровно один раз внутри NewCoupon, а затем вызывается ValidateCoupon.

	if err := ValidateCoupon(c); err != nil {
		// TODO: При ошибке NewCoupon должна возвращать Coupon{} (zero value), а не частично заполненную структуру.
		return Coupon{}, err
	}
	return c, nil
}

func main() {
	var codeInput string
	var percentInput int

	fmt.Fscan(os.Stdin, &codeInput, &percentInput)

	c, err := NewCoupon(codeInput, percentInput)
	if err != nil {
		fmt.Printf("INVALID: %s", err.Error())
		return
	}

	fmt.Printf("CREATED: %s %d", c.Code, c.Percent)
}
