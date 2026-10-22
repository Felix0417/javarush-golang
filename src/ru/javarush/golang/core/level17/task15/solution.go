package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Sentinel-ошибки строго через errors.New (по условию).
var (
	ErrBadFormat     = errors.New("bad format")
	ErrNameRequired  = errors.New("name required")
	ErrQtyOutOfRange = errors.New("qty out of range")
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	// TODO: Соберите ошибки по всем записям, не останавливаясь на первой.
	// TODO: Не добавляйте nil в список ошибок, сохраните порядок ошибок как во входе.
	// TODO: Для каждой ошибки добавьте контекст номера записи через fmt.Errorf("record %d: %w", i+1, err).
	// TODO: В конце, если ошибок нет — напечатайте ровно "IMPORTED", иначе напечатайте одну объединённую ошибку через errors.Join(errs...).
	var errs []error

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)

		err := parseAndValidate(s)
		if err != nil {
			errs = append(errs, fmt.Errorf("record %d: %w", i+1, err))
		}
	}

	if errs != nil {
		fmt.Println(errors.Join(errs...))
		return
	}

	fmt.Print("IMPORTED")
}

func parseAndValidate(s string) error {
	// TODO: Реализуйте разбор записи формата "name;qty" строго через strings.Cut.
	// TODO: Реализуйте валидацию name (не должен быть пустым).
	// TODO: Реализуйте парсинг qty строго через strconv.Atoi.
	// TODO: При ошибке парсинга qty верните ошибку с контекстом fmt.Errorf("qty not int: %w", err).
	// TODO: Проверьте диапазон qty 1..100 (включительно) и верните ErrQtyOutOfRange при нарушении.

	name, qtyStr, ok := strings.Cut(s, ";")
	if !ok {
		return ErrBadFormat
	}
	if name == "" {
		return ErrNameRequired
	}

	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		return fmt.Errorf("qty not int: %w", err)
	}
	if qty < 1 || qty > 100 {
		return ErrQtyOutOfRange
	}

	return nil
}
