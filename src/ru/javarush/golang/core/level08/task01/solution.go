package main

import "fmt"

// В этой задаче важно увидеть разницу между typed и untyped константами
// через печать значения и реального типа.

// Одиночный const (по требованиям)
const n = 5

// const-блок (по требованиям)
const (
	f = 1.25
	s = "go"
	b = true
	r = 'A'
)

// typed-константы с явным типом (по требованиям)
const (
	typedN int    = 5
	typedS string = "go"
)

func main() {
	// TODO: Выведите ровно 7 строк — по одной на каждую константу: n, f, s, b, r, typedN, typedS.
	// TODO: Формат каждой строки: имя=<значение> type=<тип>.
	// TODO: Для печати используйте fmt.Printf так, чтобы значение выводилось через %v, а тип — через %T.
	// TODO: Обратите внимание на r: в выводе значение должно получиться числом, а тип — int32.

	// Сейчас вывод намеренно неполный и в неверном формате.
	fmt.Printf("n=%v type=%T\n", n, n)
	fmt.Printf("f=%v type=%T\n", f, f)
	fmt.Printf("s=%v type=%T\n", s, s)
	fmt.Printf("b=%v type=%T\n", b, b)
	fmt.Printf("r=%v type=%T\n", r, r)
	fmt.Printf("typedN=%v type=%T\n", typedN, typedN)
	fmt.Printf("typedS=%v type=%T\n", typedS, typedS)
}
