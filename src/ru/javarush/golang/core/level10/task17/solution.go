package main

import "fmt"

// apply — чистая функция: не читает и не печатает, только считает новое значение.
func apply(op string, current int, x int) int {
	// TODO: Реализуйте логику операций add/sub/set.
	// TODO: При неизвестной операции верните current без изменений.
	switch op {
	case "add":
		return current + x
	case "sub":
		return current - x
	case "set":
		return x
	default:
		return current
	}
}

func main() {
	current := 0 // по требованиям — только локальное состояние, без глобальных переменных

	// TODO: Реализуйте цикл чтения пар (операция, число) через fmt.Scan.
	// TODO: Немедленно завершайте цикл при команде stop 0 (без вызова apply).
	// TODO: Для остальных команд обновляйте current через apply.
	var op string
	var x int

	for {
		_, err := fmt.Scan(&op, &x)
		if err != nil {
			return
		}

		if op == "stop" && x == 0 {
			break
		}
		current = apply(op, current, x)
	}

	fmt.Println(current)
}

// add 10
// add 2
// add 7
// add 4
// sub 1
// stop 0
