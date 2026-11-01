package main

import "fmt"

// Light — пользовательский тип на базе int.
type Light int

// Константы светофора через iota: 0, 1, 2.
const (
	LightRed Light = iota
	LightYellow
	LightGreen
)

// String реализует fmt.Stringer: fmt.Println(light) печатает человекочитаемо.
func (l Light) String() string {
	// TODO: Реализуйте корректное строковое представление для Light:
	// TODO: - 0 -> "red", 1 -> "yellow", 2 -> "green"
	// TODO: - для любых других значений: fmt.Sprintf("Light(%d)", int(l))
	// TODO: Используйте receiver по значению и не меняйте сигнатуру метода.
	switch l {
	case LightRed:
		return "red"
	case LightYellow:
		return "yellow"
	case LightGreen:
		return "green"
	default:
		return fmt.Sprintf("Light(%d)", int(l))
	}
}

func main() {
	var code int
	fmt.Scan(&code)

	light := Light(code)
	fmt.Println(light) // важно: без ручного вызова String()
}
