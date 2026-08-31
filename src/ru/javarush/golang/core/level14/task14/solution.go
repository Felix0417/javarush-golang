package main

import (
	"fmt"
	"strings"
)

func main() {
	var palette string
	fmt.Scan(&palette)

	// TODO: Разбейте строку palette на части через strings.Split с разделителем ":" (пустые элементы должны сохраняться).
	parts := strings.Split(palette, ":")

	// TODO: Выведите количество элементов (len(parts)) перед выводом самих элементов.
	fmt.Println(len(parts))

	// По требованию — цикл именно по индексам
	for i := 0; i < len(parts); i++ {
		// TODO: Выведите каждый элемент с новой строки в формате "i) значение", где i начинается с 1.
		fmt.Printf("%d) %s\n", i+1, parts[i])
	}
}
