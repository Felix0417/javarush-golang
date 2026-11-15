package main

/*
Вы пишете калькулятор для админки. Деление — операция простая, но при b == 0 хочется иметь причину как typed-ошибку
(чтобы можно было достать поля), а сверху добавить контекст “какая операция упала”. Для этого вы оборачиваете причину
в OpError с Unwrap, чтобы цепочка ошибок оставалась прозрачной.

Напишите программу, которая читает из stdin два целых числа dividend и divider и выполняет целочисленное деление dividend / divider.
- Если divider равно 0, функция Divide(a, b int) (int, error) должна вернуть typed error DivisionByZeroError с полем A (делимое)
и сообщением вида division by zero (a=<A>).
- Если divider не равно 0, функция возвращает результат и nil.

Ошибку из Divide нужно оборачивать в OpError (поля Op string и Err error), чтобы итоговая ошибка выглядела как <Op>: <причина>.
OpError должен иметь методы Error() и Unwrap() error, чтобы была доступна причина через errors.Unwrap.

В main при ошибке программа должна вывести сначала полную ошибку, а затем построчно все причины, двигаясь по цепочке через
errors.Unwrap.
*/

import (
	"errors"
	"fmt"
	"strconv"
)

type DivisionByZeroError struct {
	A int
}

func (e DivisionByZeroError) Error() string {
	// TODO: Верните текст typed-ошибки деления на ноль в строгом формате из условия (включая A)
	return fmt.Sprintf("division by zero (a=%s)", strconv.Itoa(e.A))
}

type OpError struct {
	Op  string
	Err error
}

func (e OpError) Error() string {
	// TODO: Верните строку ошибки верхнего уровня как описание операции + причина (ошибка Err)
	if e.Err == nil {
		return e.Op
	}
	return fmt.Sprintf("%s: %s", e.Op, e.Err)
}

func (e OpError) Unwrap() error {
	// TODO: Верните исходную причину (Err), чтобы цепочка ошибок работала через errors.Unwrap
	return e.Err
}

func Divide(a, b int) (int, error) {
	// TODO: Реализуйте целочисленное деление a / b и возврат typed-ошибки при b == 0
	if b == 0 {
		return 0, DivisionByZeroError{A: a}
	}

	return a / b, nil
}

func main() {
	var dividend, divider int
	fmt.Scan(&dividend, &divider)

	res, err := Divide(dividend, divider)
	if err == nil {
		fmt.Println(res)
		return
	}

	err = OpError{
		Op:  "divide",
		Err: err,
	}

	fmt.Println(err)

	for cause := errors.Unwrap(err); cause != nil; cause = errors.Unwrap(cause) {
		fmt.Println(cause)
	}
}
