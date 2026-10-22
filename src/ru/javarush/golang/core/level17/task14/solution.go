package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	// Sentinel-ошибки (важно: через errors.New и точные тексты)
	ErrTooShort = errors.New("too short")
	ErrNoDigit  = errors.New("no digit")
	ErrNoUpper  = errors.New("no upper")
)

func validatePassword(pw string) error {
	var errs []error

	if len(pw) < 8 {
		errs = append(errs, ErrTooShort)
	}

	hasDigit := false
	hasUpper := false

	for _, ch := range pw {
		// TODO: Реализуйте проверку символа и выставляйте hasDigit/hasUpper при обходе строки.
		if ch >= '0' && ch <= '9' {
			hasDigit = true
			continue
		}
		if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		}
		if hasDigit && hasUpper {
			break
		}
	}

	if !hasDigit {
		errs = append(errs, ErrNoDigit)
	}
	if !hasUpper {
		errs = append(errs, ErrNoUpper)
	}

	if len(errs) == 0 {
		return nil
	}

	// TODO: Верните одну ошибку, которая содержит все причины, используя errors.Join(errs...).
	return errors.Join(errs...)
}

func main() {
	var pw string
	fmt.Fscan(os.Stdin, &pw)

	err := validatePassword(pw)

	if err == nil {
		fmt.Println("OK")
	} else {
		fmt.Println(err)
	}

	// Диагностика причин — только через errors.Is
	fmt.Printf("too_short=%v\n", errors.Is(err, ErrTooShort))
	fmt.Printf("no_digit=%v\n", errors.Is(err, ErrNoDigit))
	fmt.Printf("no_upper=%v\n", errors.Is(err, ErrNoUpper))
}
