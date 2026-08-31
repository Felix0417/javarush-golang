package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	// report объявляем ДО if/else, чтобы не "потерять" значение из-за := в ветках
	var report string

	// TODO: Сформируйте строку отчёта в report:
	if n%2 == 0 {
		// TODO: Сформируйте корректный отчёт для чётного числа.
		report = "even:" + strconv.Itoa(n)
	} else {
		// TODO: Сформируйте корректный отчёт для нечётного числа.
		report = "odd:" + strconv.Itoa(n)
	}

	fmt.Print(report) // выводим ровно одну строку без лишних сообщений
}
