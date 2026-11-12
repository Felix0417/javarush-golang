package main

/*
Вы пишете маленький “логгер координат” для игры: ему всё равно, что именно вы передали — главное, чтобы объект умел
превращать себя в строку. Именно поэтому вы опираетесь на минимальный контракт fmt.Stringer, а не на конкретный тип точки.

Создайте структуру Point с полями X и Y типа int. Реализуйте для неё метод String() string, чтобы строка выглядела
как Point(X,Y) без пробелов, например Point(2,3).

Создайте функцию PrintOne(v fmt.Stringer), которая печатает ровно одну строку: результат v.String().

В main вызовите PrintOne для Point(2,3) и Point(0,0) — программа должна вывести две строки в этом порядке.
Point(2,3)
Point(0,0)
*/

import "fmt"

type Point struct {
	X, Y int
}

// Value receiver: чтобы и Point, и *Point удовлетворяли fmt.Stringer.
func (p Point) String() string {
	// TODO: Реализуйте строковое представление точки в формате Point(X,Y) без пробелов.
	// TODO: Убедитесь, что для Point{2,3} получается ровно "Point(2,3)".
	return fmt.Sprintf("Point(%d,%d)", p.X, p.Y)
}

// Принимаем минимальный контракт: нам нужна только строковая форма.
func PrintOne(v fmt.Stringer) {
	// TODO: Напечатайте ровно одну строку: результат v.String().
	fmt.Println(v.String())
}

func main() {
	PrintOne(Point{2, 3})
	PrintOne(Point{0, 0})
}
