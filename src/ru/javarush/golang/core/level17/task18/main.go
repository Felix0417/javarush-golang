package main

import (
	"errors"
	"fmt"

	"myproject/apperr"
)

func main() {
	var mode int
	fmt.Scan(&mode)

	var err error
	switch mode {
	case 0:
		err = nil
	case 1:
		// TODO: Сформируйте ошибку для режима 1 (wrapping ErrValidation) по условию, используя fmt.Errorf и %w.
		err = fmt.Errorf("bad input: %w", apperr.ErrValidation)
	case 2:
		// TODO: Сформируйте ошибку для режима 2 (wrapping ErrNotFound) по условию, используя fmt.Errorf и %w.
		err = fmt.Errorf("missing: %w", apperr.ErrNotFound)
	case 3:
		// TODO: Сформируйте ошибку для режима 3 по условию: внутри должна быть объединённая ошибка (join), а затем wrapping через fmt.Errorf и %w.
		err = fmt.Errorf("read: %w", errors.Join(apperr.ErrIO, errors.New("unexpected EOF")))
	case 4:
		// TODO: Сформируйте ошибку для режима 4 (wrapping ErrInternal) по условию, используя fmt.Errorf и %w.
		err = fmt.Errorf("bug: %w", apperr.ErrInternal)
	default:
		err = apperr.ErrInternal
	}

	fmt.Println(apperr.UserMessage(err))
}
