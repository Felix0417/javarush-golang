package main

import "fmt"

func main() {
	var (
		x        int
		sum      int
		max      int
		posCount int
		hasAny   bool
	)

	for {
		if _, err := fmt.Scan(&x); err != nil {
			break
		}

		sum += x

		// TODO: Увеличивать счётчик только для чисел, которые строго больше 0 (x > 0).
		if x > 0 {
			posCount++
		}

		// TODO: Реализовать корректное вычисление максимума, включая случай,
		// когда все числа отрицательные (нельзя полагаться на max=0).
		if x > max {
			max = x
		}

		hasAny = true
	}

	if !hasAny {
		fmt.Println(0)
		fmt.Println(0)
		fmt.Println(0)
		return
	}

	fmt.Println(sum)
	fmt.Println(max)
	fmt.Println(posCount)
}
