package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	var xCoord, yCoord int
	fmt.Scan(&xCoord, &yCoord)

	// TODO: Создайте значение p типа Point с помощью позиционного литерала (без указания имён полей),
	// TODO: используя введённые xCoord и yCoord.
	p := Point{xCoord, yCoord}

	// TODO: Выведите координаты ровно одной строкой в формате: X=<...> Y=<...> (между частями один пробел).
	fmt.Printf("X=%d Y=%d", p.X, p.Y)
}
