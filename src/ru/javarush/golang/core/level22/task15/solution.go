package main

import "fmt"

type Point struct {
	X int
	Y int
}

// String реализован только для *Point, чтобы отличать вывод значения и указателя.
func (p *Point) String() string {
	// TODO: Добавьте проверку p == nil и верните безопасную строку согласно условию.
	// TODO: Для ненулевого указателя верните координаты в строгом формате согласно условию (без пробелов).
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var p Point
	p.X = x
	p.Y = y

	fmt.Println(p)
	fmt.Println(&p)
}
