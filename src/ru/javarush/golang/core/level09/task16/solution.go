package main

import (
	"errors"
	"fmt"
	"os"
)

func minMax(nums ...int) (min int, max int, err error) {
	if len(nums) == 0 {
		// Нельзя вычислить min/max "из воздуха".
		return 0, 0, errors.New("no numbers")
	}

	min, max = nums[0], nums[0]

	// TODO: Реализуйте поиск минимума и максимума за один проход по nums (без сортировки и без дополнительных проходов)
	for _, num := range nums {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}

	return min, max, nil
}

func formatStats(name string, nums ...int) (string, error) {
	// Важно: прокидываем variadic как minMax(nums...) (с тремя точками).
	min, max, err := minMax(nums...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s: min=%d max=%d count=%d", name, min, max, len(nums)), nil
}

func main() {
	in := os.Stdin

	var name string
	var count int
	fmt.Fscan(in, &name, &count)

	var a1, a2, a3, a4, a5 int
	if count >= 1 {
		fmt.Fscan(in, &a1)
	}
	if count >= 2 {
		fmt.Fscan(in, &a2)
	}
	if count >= 3 {
		fmt.Fscan(in, &a3)
	}
	if count >= 4 {
		fmt.Fscan(in, &a4)
	}
	if count >= 5 {
		fmt.Fscan(in, &a5)
	}

	// По требованию: вызываем formatStats с 0..5 аргументами через ветвление.
	var s string
	var err error
	switch count {
	case 0:
		s, err = formatStats(name)
	case 1:
		s, err = formatStats(name, a1)
	case 2:
		s, err = formatStats(name, a1, a2)
	case 3:
		s, err = formatStats(name, a1, a2, a3)
	case 4:
		s, err = formatStats(name, a1, a2, a3, a4)
	default: // 5
		s, err = formatStats(name, a1, a2, a3, a4, a5)
	}

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	fmt.Println(s)
}
