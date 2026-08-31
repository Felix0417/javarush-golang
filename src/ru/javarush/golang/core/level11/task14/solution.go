package main

import "fmt"

func addOne(s []int, x int) []int {
	// TODO: Добавьте число x в конец слайса s через append и верните обновлённый слайс.
	// TODO: Важно: результат append нельзя игнорировать — его нужно сохранить обратно в переменную слайса.
	s = append(s, x)
	return s
}

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	var extraValue int
	fmt.Scan(&extraValue)

	// TODO: Вызовите addOne и сохраните результат обратно в numbers (результат функции нельзя игнорировать).
	numbers = addOne(numbers, extraValue)

	for i, v := range numbers {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}
