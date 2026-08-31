package main

import "fmt"

func main() {
	var readings []int // важно: это nil-слайс, не используем make

	for {
		var x int
		fmt.Scan(&x) // читаем поток целых чисел до завершающего 0

		if x == 0 {
			break // 0 — сигнал конца, его не сохраняем
		}

		// TODO: Добавьте считанное число x в слайс readings через append с сохранением результата.
		readings = append(readings, x)
	}

	// TODO: Выведите все значения из readings в исходном порядке в одну строку через пробел.
	// TODO: Убедитесь, что при пустом readings (например, если первым числом был 0) выводится пустая строка.

	for i, v := range readings {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", v)
	}
	fmt.Println()
}
