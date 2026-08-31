package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	// По требованию: создаём слайс длины N через make и заполняем из ввода.
	tasks := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&tasks[i])
	}

	var holeIndex int
	fmt.Scan(&holeIndex)

	// TODO: Проверьте, что holeIndex находится в диапазоне 0 ≤ holeIndex < N-1 (при N=0 или N=1 сдвиг не выполняется).
	// TODO: Если индекс корректный, выполните сдвиг элементов влево РОВНО одной операцией copy, не меняя длину tasks.
	if holeIndex >= 0 && holeIndex < N-1 {
		copy(tasks[holeIndex:], tasks[holeIndex+1:])
	}

	// Ровно N чисел через пробел.
	for i := 0; i < N; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(tasks[i])
	}
	fmt.Println()
}
