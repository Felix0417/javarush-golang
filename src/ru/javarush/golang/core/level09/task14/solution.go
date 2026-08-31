package main

import (
	"fmt"
	"os"
)

func sum(nums ...int) int {
	// TODO: Реализуйте подсчёт суммы всех чисел из nums.
	// TODO: Пустой вызов sum() должен возвращать 0.
	var sum int
	for _, s := range nums {
		sum += s
	}
	return sum
}

func printSum(prefix string, nums ...int) {
	// TODO: Напечатайте ровно одну строку в формате: "<prefix> <sum>\n".
	// TODO: Сумму нужно получать только вызовом sum(nums...) (с прокидыванием nums...).

	fmt.Printf("%s %d\n", prefix, sum(nums...))
}

func main() {
	var prefix string
	var count int
	fmt.Fscan(os.Stdin, &prefix, &count)

	switch count {
	case 0:
		printSum(prefix)
	case 1:
		var a int
		fmt.Fscan(os.Stdin, &a)
		printSum(prefix, a)
	case 2:
		var a, b int
		fmt.Fscan(os.Stdin, &a, &b)
		printSum(prefix, a, b)
	case 3:
		var a, b, c int
		fmt.Fscan(os.Stdin, &a, &b, &c)
		printSum(prefix, a, b, c)
	}
}
