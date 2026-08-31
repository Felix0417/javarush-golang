package main

import "fmt"

// abs возвращает модуль числа через именованный результат res.
// Важно: используйте if и "пустой return", но при этом res должен быть явно присвоен во всех ветках.
func abs(x int) (res int) {
	if x < 0 {
		res = -x
		return
	}

	// TODO: Присвойте res корректное значение для случая x >= 0 (иначе res останется равным zero value).
	res = x
	return
}

func main() {
	var positionX int
	fmt.Scan(&positionX)

	fmt.Println(abs(positionX))
}
