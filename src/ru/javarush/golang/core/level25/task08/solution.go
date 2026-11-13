package main

/*
Вы делаете имитацию фоновой работы: иногда всё прошло успешно, иногда ошибка временная (можно попробовать ещё раз), а
иногда ошибка фатальная (повтор бесполезен). В Go это часто делают так: проверяют, реализует ли ошибка дополнительный
интерфейс, например Temporary() bool.

Из stdin читаются code (int) и msg (строка без пробелов). Объявите интерфейс temporary, который встраивает error и добавляет
метод Temporary() bool. Затем объявите два типа ошибок: TempError (реализует error и Temporary() bool, возвращающий true)
и FailError (реализует только error).

Реализуйте doWork(code int, msg string) error: при code == 0 вернуть nil, при code == 1 вернуть TempError с сообщением msg,
при code == 2 вернуть FailError с сообщением msg.

Затем реализуйте handle(err error) string, которая делает безопасный type assertion err.(temporary). Если ошибка temporary
и Temporary() возвращает true, результат должен быть retry: <msg>, иначе fatal: <msg>. Для nil верните ok. Напечатайте
результат handle одной строкой.
*/

import (
	"fmt"
	"os"
)

// temporary — "расширенная" ошибка: обычный error + признак временности.
type temporary interface {
	error
	Temporary() bool
}

type TempError struct {
	msg string
}

func (e TempError) Error() string { return e.msg }

func (e TempError) Temporary() bool {
	// TODO: Реализовать признак временности для временной ошибки (должен отличаться от фатальной).
	return true
}

type FailError struct {
	msg string
}

func (e FailError) Error() string { return e.msg }

func doWork(code int, msg string) error {
	// TODO: Реализовать выбор типа ошибки по code:
	// TODO: 0 -> nil, 1 -> TempError, 2 -> FailError.
	switch code {
	case 0:
		return nil
	case 1:
		return TempError{msg: msg}
	case 2:
		return FailError{msg: msg}
	default:
		return fmt.Errorf("unknown code: %d", code)
	}
}

func handle(err error) string {
	if err == nil {
		return "ok"
	}

	// Требование задачи: использовать безопасную форму type assertion "t, ok :=".
	t, ok := err.(temporary)

	// TODO: Реализовать обработку временных/фатальных ошибок через интерфейс temporary:
	// TODO: если ok == true и t.Temporary() == true -> "retry: <msg>", иначе -> "fatal: <msg>".
	if ok && t.Temporary() {
		return fmt.Sprintf("retry: %s", err.Error())
	}
	return "fatal: " + err.Error()
}

func main() {
	var code int
	var msg string
	fmt.Fscan(os.Stdin, &code, &msg)

	fmt.Println(handle(doWork(code, msg)))
}
