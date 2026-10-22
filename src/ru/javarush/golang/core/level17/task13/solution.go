package main

import (
	"errors"
	"fmt"
)

// Sentinel-ошибки из условия.
var (
	ErrTitleRequired      = errors.New("title required")
	ErrPriorityOutOfRange = errors.New("priority out of range")
)

func main() {
	var title string
	var priority int
	fmt.Scan(&title, &priority)

	// Собираем нарушения в errs.
	var errs []error

	// TODO: Реализовать независимую валидацию title и priority (без раннего выхода),
	// TODO: добавить в errs только ненулевые ошибки, сохраняя порядок проверок.
	if title == "-" {
		errs = append(errs, ErrTitleRequired)
	}
	if priority < 1 || priority > 5 {
		errs = append(errs, ErrPriorityOutOfRange)
	}

	// TODO: Объединить ошибки через errors.Join(errs...) и использовать результат
	// TODO: для корректного вывода: "OK" при nil, иначе печать объединённой ошибки.
	err := errors.Join(errs...)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
}
