package main

import "fmt"

const (
	// Три флага как разные биты: 001, 010, 100
	R uint = 1 << iota
	W
	X
)

func main() {
	var raw uint
	var op string
	var flag string

	fmt.Scan(&raw, &op, &flag)

	var mask uint
	switch flag {
	case "R":
		mask = R
	case "W":
		mask = W
	case "X":
		mask = X
	default:
		fmt.Print("error")
		return
	}

	// mask используется, чтобы шаблон компилировался до реализации операций.
	_ = mask

	switch op {
	case "set":
		// TODO: Реализуйте операцию set: установить бит флага в raw.
		raw = raw | mask
		fmt.Printf("%d %03b", raw, raw)
	case "clear":
		// TODO: Реализуйте операцию clear: очистить бит флага в raw.
		raw = raw &^ mask
		fmt.Printf("%d %03b", raw, raw)
	case "toggle":
		// TODO: Реализуйте операцию toggle: инвертировать бит флага в raw.
		raw = raw ^ mask
		fmt.Printf("%d %03b", raw, raw)
	case "check":
		// TODO: Реализуйте операцию check: проверить, установлен ли бит флага в raw.
		if raw&mask != 0 {
			fmt.Print("on")
		} else {
			fmt.Print("off")
		}
	default:
		fmt.Print("error")
	}
}
