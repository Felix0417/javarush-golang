package main

import "fmt"

func main() {
	original := []string{"red", "green", "blue"}

	alias := original

	// TODO: Замените второй элемент палитры только через alias (без прямого изменения original)
	alias[1] = "YELLOW"

	fmt.Println(alias)
	fmt.Println(original)

	// TODO: Выведите булево значение строго из выражения сравнения адресов первых элементов
	fmt.Println(&original[0] == &alias[0])
}
