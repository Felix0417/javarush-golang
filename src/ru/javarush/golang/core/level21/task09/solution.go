package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	// Создаём именно значения struct (не указатели) — такие структуры сравнимы через ==
	p1 := Point{X: x1, Y: y1}
	p2 := Point{X: x2, Y: y2}

	// TODO: Сравните p1 и p2 ТОЛЬКО оператором == и выведите:
	// TODO: EQUAL, если точки совпадают, иначе NOT EQUAL.
	// TODO: Нельзя сравнивать поля вручную (p1.X == p2.X ...), по условию требуется == для struct.
	if p1 == p2 {
		fmt.Println("EQUAL")
	} else {
		fmt.Println("NOT EQUAL")
	}
}
