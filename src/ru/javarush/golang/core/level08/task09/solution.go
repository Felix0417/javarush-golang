package main

import "fmt"

type TrafficLight int

const (
	Red TrafficLight = iota
	Yellow
	Green
)

// TODO: Объявите собственный тип TrafficLight на базе int (если он ещё не объявлен выше).
// TODO: Объявите ровно один const-блок, в котором константы Red, Yellow, Green задаются через iota.
// TODO: Не присваивайте константам Red, Yellow, Green числовые значения вручную.

func main() {
	// TODO: Выведите значения Red, Yellow и Green в одной строке: три числа, разделённые одиночными пробелами.
	// TODO: Для вывода используйте пакет fmt.
	fmt.Printf("%d %d %d", Red, Yellow, Green)
}
