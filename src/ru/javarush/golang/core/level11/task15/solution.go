package main

import "fmt"

func main() {
	// Исходный слайс. Важно: программа без ввода.
	s := []int{1, 2}

	addWrong(s)
	fmt.Println(s) // должно остаться: [1 2]

	s = addRight(s)
	fmt.Println(s) // должно стать: [1 2 999]
}

// append возвращает новый (возможно) слайс-хедер, но здесь мы его "теряем" — это и есть ошибка.
func addWrong(s []int) {
	append(s, 999)
}

func addRight(s []int) []int {
	// TODO: Реализуйте корректное добавление 999 через append: сохраните результат append и верните обновлённый слайс.
	s = append(s, 999)
	return s
}
