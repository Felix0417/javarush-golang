package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Print("0 -1")
		return
	}

	var max int
	fmt.Scan(&max)
	maxIdx := 0

	// Читаем оставшиеся значения "на лету", без слайсов/массивов.
	for i := 1; i < n; i++ {
		var x int
		fmt.Scan(&x)

		if x > max {
			max = x
			maxIdx = i
		}

		// TODO: Реализуйте обновление рекорда и индекса его первого появления.
		// TODO: Индекс обновляйте только если найдено значение СТРОГО больше текущего максимума.
		// TODO: Хранить все замеры в слайсе/массиве нельзя — обрабатывайте значения при чтении.
	}

	fmt.Printf("%d %d\n", max, maxIdx)
}
