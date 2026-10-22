package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// По условию: слайс строго фиксированной длины через make([]int, n).
	nums := make([]int, n)

	// Сумма — отдельная переменная типа int.
	sum := 0

	// TODO: считайте ровно n целых чисел и заполните nums по индексам (без append)
	// TODO: посчитайте сумму всех элементов nums в переменной sum (тип int)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)

		// TODO: запишите считанное значение v в слайс nums по индексу i
		// TODO: обновите сумму sum, добавив сохранённое значение
		nums[i] = v
		sum += v
	}

	// "Паспорт" хранилища: len, cap и сумма.
	fmt.Printf("%d %d %d\n", len(nums), cap(nums), sum)
}
