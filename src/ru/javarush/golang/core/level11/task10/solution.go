package main

import "fmt"

func main() {
	var mode string
	fmt.Scan(&mode)

	// s должен быть []int в одном из трёх "режимов пустоты"
	var s []int

	switch mode {
	case "nil":
		// TODO: Реализуйте создание nil-слайса (в этом режиме s должен остаться nil).
	case "empty":
		// TODO: Реализуйте создание пустого, но НЕ nil, слайса.
		s = []int{}
	case "cap":
		var k int
		fmt.Scan(&k)

		// TODO: Реализуйте создание пустого слайса с ёмкостью k.
		s = make([]int, 0, k)
	}

	fmt.Printf("nil=%v len=%d cap=%d\n", s == nil, len(s), cap(s))
}
