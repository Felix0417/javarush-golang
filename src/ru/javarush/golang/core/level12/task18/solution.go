package main

import "fmt"

func main() {
	// Базовый слайс: len=2 cap=2 — места для роста нет.
	base := make([]int, 2, 2)

	// TODO: Запишите в base значения 10 и 20 по индексам 0 и 1.
	base[0] = 10
	base[1] = 20

	// TODO: Создайте sub как полный срез base[:] до append.
	sub := base[:]

	// TODO: Выведите строку base(before) в точном формате из условия.
	fmt.Printf("base(before): len=%d cap=%d %v\n", len(base), cap(base), base)

	// TODO: Выполните расширение под-слайса строго операцией sub = append(sub, 30).
	sub = append(sub, 30)

	// TODO: Выведите строку base(after) в точном формате из условия.
	fmt.Printf("base(after): len=%d cap=%d %v\n", len(base), cap(base), base)

	// TODO: Выведите строку sub(after) в точном формате из условия (без len/cap для sub).
	fmt.Printf("sub(after): %v\n", sub)
}
