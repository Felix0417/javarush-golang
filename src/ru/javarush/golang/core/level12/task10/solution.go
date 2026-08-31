package main

import "fmt"

func main() {
	// Два независимых базовых слайса с одинаковым содержимым.
	base1 := []string{"A", "B", "C", "D"}
	base2 := []string{"A", "B", "C", "D"}

	// TODO: Возьмите под-слайс на первые 2 элемента из base1 (без ограничения capacity).
	// TODO: Добавьте строку "X" через append, сохранив результат обратно в view1.
	view1 := base1[:2]
	view1 = append(view1, "X")

	fmt.Println(base1)

	// TODO: Возьмите под-слайс на первые 2 элемента из base2, но с ограничением capacity (full slice expression).
	// TODO: Добавьте строку "X" через append, сохранив результат обратно в view2.
	view2 := base2[:2:2]
	view2 = append(view2, "X")

	fmt.Println(base2)
}
