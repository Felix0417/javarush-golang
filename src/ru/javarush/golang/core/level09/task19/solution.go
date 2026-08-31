package main

import "fmt"

// makeStepper возвращает функцию-счётчик.
//
// ВАЖНО: состояние счётчика должно храниться внутри замыкания.
func makeStepper(start, step int) func() int {
	// TODO: Реализуйте хранение текущего значения внутри замыкания.
	// TODO: При каждом вызове возвращайте очередное значение, начиная со start,
	// TODO: и изменяйте внутреннее состояние на step для следующего вызова.

	current := start

	return func() int {
		// TODO: Верните текущее значение счётчика и обновите его для следующего тика.
		old := current
		current += step
		return old
	}
}

func main() {
	var startValue, stepSize, ticksCount int
	fmt.Scan(&startValue, &stepSize, &ticksCount)

	if ticksCount <= 0 {
		return
	}

	next := makeStepper(startValue, stepSize)
	for i := 0; i < ticksCount; i++ {
		fmt.Println(next())
	}
}
