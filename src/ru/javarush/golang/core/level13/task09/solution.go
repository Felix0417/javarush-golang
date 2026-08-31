package main

import "fmt"

func removeAtStable(s []int, removeIndex int) []int {
	// Если индекс вне диапазона — возвращаем исходный слайс без изменений.
	if removeIndex < 0 || removeIndex >= len(s) {
		return s
	}

	// TODO: Реализуйте stable-удаление элемента из слайса, сохраняя порядок.
	// TODO: Сдвиньте элементы влево с помощью copy так, чтобы элемент с removeIndex исчез.
	// TODO: Очистите освободившийся последний слот через clear (до уменьшения длины).
	// TODO: Уменьшите длину через reslice и верните новый слайс.

	copy(s[removeIndex:], s[removeIndex+1:])
	clear(s[len(s)-1:])
	return s[:len(s)-1]
}

func main() {
	numbers := []int{5, 10, 15, 20, 25}
	removeIndex := 2

	// Важно сохранить результат обратно (не через :=), иначе легко ошибиться с использованием старого слайса.
	numbers = removeAtStable(numbers, removeIndex)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
}
