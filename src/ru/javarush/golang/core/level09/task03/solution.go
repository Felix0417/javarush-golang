package main

import "fmt"

func max2(x, y int) int {
	// TODO: Реализуйте функцию, которая возвращает большее из двух целых чисел (при равенстве верните это же значение).
	if x > y {
		return x
	}
	return y
}

func max3(a, b, c int) int {
	// TODO: Реализуйте максимум из трёх чисел через вызовы max2 (без одного большого условия, сравнивающего сразу три значения).
	// Временная логика: сравниваем только первые два значения, третье игнорируется.
	return max2(max2(a, b), c)
}

func main() {
	var firstScore, secondScore, thirdScore int
	fmt.Scan(&firstScore, &secondScore, &thirdScore)

	fmt.Print(max3(firstScore, secondScore, thirdScore))
}
