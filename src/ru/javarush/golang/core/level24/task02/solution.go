package main

/*
Представьте, что вы пишете простую систему, которая складывает показания двух датчиков. Сегодня датчики очень примитивные:
каждый умеет сообщать только одно число. Но вы заранее делаете код гибким — пусть функция суммирования знает не про
конкретные структуры, а только про то, что у "источника числа" можно спросить значение.

Программа читает из stdin два целых числа. Вы объявляете интерфейс Valuer с методом Value() int, структуру NumberBox с
полем N int и методом Value() int, который просто возвращает N. Затем пишете функцию Sum(x, y Valuer) int, которая
возвращает x.Value() + y.Value(). В main вы упаковываете оба введённых числа в NumberBox, передаёте их в Sum как Valuer
и выводите сумму.
*/

import "fmt"

// Valuer — "источник числа": нам важно только уметь запросить значение.
type Valuer interface {
	Value() int
}

// NumberBox упаковывает одно число и реализует Valuer.
type NumberBox struct {
	N int
}

func (b NumberBox) Value() int {
	// TODO: Верните значение поля N, чтобы NumberBox корректно реализовывал интерфейс Valuer.
	return b.N
}

// Sum работает через интерфейс, а не через конкретные структуры.
func Sum(x, y Valuer) int {
	// TODO: Реализуйте сумму через вызовы Value() у x и y.
	return x.Value() + y.Value()
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	x := NumberBox{N: a}
	y := NumberBox{N: b}

	fmt.Println(Sum(x, y))
}
