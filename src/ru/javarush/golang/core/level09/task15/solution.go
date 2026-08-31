package main

import (
	"errors"
	"fmt"
)

func average(nums ...int) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("no numbers")
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	// TODO: Посчитайте среднее арифметическое как float64 (без целочисленного деления).
	return float64(sum) / float64(len(nums)), nil
}

func main() {
	var count int
	fmt.Scan(&count)

	// читаем ровно count чисел (0..4)
	var a, b, c, d int
	switch count {
	case 0:
		// ничего не читаем
	case 1:
		fmt.Scan(&a)
	case 2:
		fmt.Scan(&a, &b)
	case 3:
		fmt.Scan(&a, &b, &c)
	case 4:
		fmt.Scan(&a, &b, &c, &d)
	}

	// вызываем average через ветвление по count (0..4)
	var (
		avg float64
		err error
	)
	switch count {
	case 0:
		avg, err = average()
	case 1:
		avg, err = average(a)
	case 2:
		// TODO: Вызовите average с двумя аргументами.
		avg, err = average(a, b)
	case 3:
		// TODO: Вызовите average с тремя аргументами.
		avg, err = average(a, b, c)
	case 4:
		// TODO: Вызовите average с четырьмя аргументами.
		avg, err = average(a, b, c, d)
	}

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Printf("%.2f", avg)
}
