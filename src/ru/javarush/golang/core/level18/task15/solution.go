package main

import (
	"errors"
	"fmt"
	"os"
)

// panicToError превращает любое значение из panic в error.
//
// TODO: Реализуйте преобразование panic value в error по правилам задачи:
// TODO: - если r имеет тип error, нужно вернуть error с wrapping через %w
// TODO: - иначе нужно вернуть error с форматированием значения r через %v
func panicToError(r any) error {
	if e, ok := r.(error); ok {
		return fmt.Errorf("panic: %w", e)
	}
	return fmt.Errorf("panic: %v", r)
}

// run — "опасный" код: читает mode и может паниковать по условию.
func run() error {
	var mode string
	fmt.Fscan(os.Stdin, &mode)

	// TODO: Реализуйте ветки поведения по mode:
	// TODO: - при mode == "err" должна быть panic(error) с нужным текстом
	// TODO: - при mode == "str" должна быть panic(string) с нужным текстом
	// TODO: - иначе вернуть nil без паники
	if mode == "err" {
		panic(errors.New("disk failed"))
	}
	if mode == "str" {
		panic("boom")
	}
	return nil
}

// safeRun — единственная "страховочная оболочка": ловит panic через defer+recover
// и конвертирует её в error.
func safeRun() (err error) {
	defer func() {
		if r := recover(); r != nil {
			// TODO: Используйте panicToError(r) и убедитесь, что формат ошибки соответствует требованиям.
			err = panicToError(r)
		}
	}()

	return run()
}

func main() {
	if err := safeRun(); err != nil {
		fmt.Println("app error:", err)
		return
	}
	fmt.Println("ok")
}
