package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var rawTitle string

	// TODO: Считайте заголовок задачи из stdin так, чтобы не потерять пробелы по краям (нужно прочитать всю строку).
	fmt.Scanf("%[^\n]\n", &rawTitle)

	title := strings.TrimSpace(rawTitle)

	// TODO: ID должен быть строкой и вычисляться от длины нормализованного заголовка.
	id := strconv.Itoa(len(title))

	// TODO: Выведите результат в точном формате: id: <ID> title: <TITLE> (без лишних символов).
	fmt.Printf("id: %s title: %s\n", id, title)
}

// kuku kuku
