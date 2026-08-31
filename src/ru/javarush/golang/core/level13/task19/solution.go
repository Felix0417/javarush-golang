package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)

	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
	}

	var pos int
	var m int
	fmt.Scan(&pos, &m)

	toInsert := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&toInsert[i])
	}

	// TODO: Реализуйте проверку корректности pos (диапазон [0; len(words)]).
	// TODO: Если позиция корректна — вставьте все слова из toInsert начиная с pos, используя slices.Insert.
	// TODO: Не забудьте сохранить результат вставки обратно в words.
	// TODO: Если toInsert — это слайс, его нужно передать в Insert как variadic-аргументы.
	if pos >= 0 && pos <= len(words) {
		words = slices.Insert(words, pos, toInsert...)
	}
	//
	// Временное поведение шаблона: вставка не выполняется.
	words = slices.Clone(words)

	for i, w := range words {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(w)
	}
}
