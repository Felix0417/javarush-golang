package main

import (
	"errors"
	"fmt"

	"myproject/app"
)

func main() {
	// 1) nil
	var e0 error = nil

	// TODO: Соберите ошибку класса validation через fmt.Errorf(... %w ...) вокруг app.ErrValidation.
	e1 := fmt.Errorf("bad input: %w", app.ErrValidation)

	// TODO: Соберите ошибку класса notfound через fmt.Errorf(... %w ...) вокруг app.ErrNotFound.
	e2 := fmt.Errorf("entity lookup: %w", app.ErrNotFound)

	// TODO: Соберите ошибку класса io: errors.Join(app.ErrIO, cause) и затем оберните контекстом через fmt.Errorf(... %w ...).
	cause := errors.New("disk timeout")
	e3 := fmt.Errorf("read file: %w", errors.Join(app.ErrIO, cause))

	// TODO: Соберите ошибку класса internal через fmt.Errorf(... %w ...) вокруг app.ErrInternal.
	e4 := fmt.Errorf("bug: %w", app.ErrInternal)

	for _, err := range []error{e0, e1, e2, e3, e4} {
		fmt.Println(app.Class(err))
	}
}
