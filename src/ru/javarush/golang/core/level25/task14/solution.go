package main

/*
Вы настраиваете "коридор безопасности" для датчика: он выдаёт число x, и оно считается нормальным только если попадает в
диапазон от min до max. Если x вылетает за границы, вы хотите вернуть typed-ошибку с доступом к полям, чтобы на границе
программы собрать стабильное сообщение без разборов текста.

Напишите программу, которая читает из stdin три целых числа: minAllowed, maxAllowed и measuredValue. Число считается
корректным, если minAllowed ≤ measuredValue ≤ maxAllowed.

Реализуйте функцию ValidateRange(min, max, x int) error. Если x в диапазоне, функция должна вернуть nil. Если x вне
диапазона, функция должна вернуть typed error RangeError с полями Min, Max и Value.

В main программа должна печатать OK, если ошибки нет. Если ошибка есть и это RangeError, программа должна печатать
строку OUT OF RANGE: value=<Value> expected=<Min>..<Max>, причём значения берутся из полей ошибки, а не “вырезаются”
из текста Error().
*/

import (
	"fmt"
	"os"
)

type RangeError struct {
	Min   int
	Max   int
	Value int
}

func (e RangeError) Error() string {
	return fmt.Sprintf("value %d out of range %d..%d", e.Value, e.Min, e.Max)
}

func ValidateRange(min, max, x int) error {
	// TODO: Реализуйте проверку, что x находится в диапазоне [min..max] (границы включительно).
	// TODO: Если x в диапазоне — верните nil.
	// TODO: Если x вне диапазона — верните typed error RangeError с заполненными полями Min/Max/Value.
	if x >= min && x <= max {
		return nil
	}
	return RangeError{Min: min, Max: max, Value: x}
}

func main() {
	var minAllowed, maxAllowed, measuredValue int
	fmt.Fscan(os.Stdin, &minAllowed, &maxAllowed, &measuredValue)

	err := ValidateRange(minAllowed, maxAllowed, measuredValue)
	if err == nil {
		fmt.Println("OK")
		return
	}

	switch e := err.(type) {
	case RangeError:
		// TODO: Сформируйте вывод строго в формате:
		// TODO: OUT OF RANGE: value=<Value> expected=<Min>..<Max>
		// TODO: Важно: числа нужно брать из полей ошибки, без разбора/парсинга текста Error().
		fmt.Printf("OUT OF RANGE: value=%d expected=%d..%d\n", e.Value, e.Min, e.Max)
	default:
		fmt.Println(err)
	}
}
