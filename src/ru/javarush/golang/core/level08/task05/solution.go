package main

import "fmt"

// Level — отдельный тип для "уровня", чтобы не путать с обычным int.
type Level int

func main() {
	level := Level(3)

	// TODO: Выведите первую строку в точном формате из условия:
	// TODO: level=<значение> type=<тип>
	fmt.Printf("level=%v type=%T\n", level, level)

	// TODO: Выполните явное преобразование Level -> int (через int(level))
	// TODO: и выведите вторую строку в точном формате из условия:
	// TODO: asInt=<значение> type=<тип>
	asInt := int(level)
	fmt.Printf("asInt=%v type=%T\n", asInt, asInt)
}
