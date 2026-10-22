package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type panicError struct {
	v     any
	stack []byte
}

func (e *panicError) Error() string {
	return fmt.Sprintf("panic: %v", e.v)
}

// run специально без проверок: a / b должен реально паниковать при b == 0.
func run(a, b int) int {
	return a / b
}

// safeRun — граница приложения. Здесь должна быть вся защита от паники.
// Сейчас функция реализована частично.
func safeRun(a, b int) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			// TODO: Перехватите панику и верните error так, чтобы наверху можно было:
			// TODO: 1) получить текст паники (panic: <...>)
			// TODO: 2) при debug-режиме вывести stack trace, снятый в момент паники
			//
			// Важно: обработка паники должна быть только здесь, а не в run и не в main.
			err = &panicError{
				v:     r,
				stack: debug.Stack(),
			}
		}
	}()

	res = run(a, b)
	return res, nil
}

func main() {
	var debugMode string
	var a, b int
	fmt.Scan(&debugMode, &a, &b)

	res, err := safeRun(a, b)
	if err != nil {
		if debugMode == "debug" {
			// TODO: Если был recover от паники, выведите:
			// TODO: panic stack:
			// TODO: <stack trace>
			// TODO: Stack trace должен быть получен через runtime/debug.Stack().
			fmt.Println("panic stack:")
			if pe, ok := errors.AsType[*panicError](err); ok {
				fmt.Printf("%s", pe.stack)
			}
		}
		fmt.Printf("app error: %s\n", err.Error())
		return
	}

	fmt.Printf("result: %d\n", res)
}
