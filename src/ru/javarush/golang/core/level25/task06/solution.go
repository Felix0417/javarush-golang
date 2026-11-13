package main

/*
Вы собираете консольный вывод для отладочного режима. Если объект умеет красиво печататься сам (реализует fmt.Stringer ),
вы хотите довериться ему. Если нет — показать тип и значение “как есть”, чтобы разработчик сразу видел, что прилетело.

Программа должна объявить структуру Point с полями X и Y типа int и методом String() string , который возвращает строку
вида "Point(X=3,Y=7)" .

Затем программа должна прочитать из stdin три целых числа: x , y и z .

Далее создайте две переменные типа any : firstItem хранит Point{X: x, Y: y} , а secondItem хранит число z (как int ).
Реализуйте функцию render(x any) , которая делает безопасный type assertion к интерфейсу fmt.Stringer : если получилось —
печатает v.String() , иначе — печатает строку формата "type=<T> value=<v>" (через "%T" и "%v" ). В main вызовите render
для обоих значений.
*/

import (
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	// TODO: Реализуйте строковое представление Point строго в формате Point(X=<x>,Y=<y>)
	return fmt.Sprintf("Point(X=%d,Y=%d)", p.X, p.Y)
}

func render(x any) {
	// TODO: Выполните безопасный type assertion к fmt.Stringer ровно в форме: v, ok := x.(fmt.Stringer)
	v, ok := x.(fmt.Stringer)
	// TODO: Если ok=true, напечатайте v.String() (одна строка вывода)
	if ok {
		fmt.Println(v.String())
		return
	}

	// TODO: Если ok=false, напечатайте fallback в формате: type=<T> value=<v> (через %T и %v)
	fmt.Printf("type=%T value=%v\n", x, x)
}

func main() {
	var x, y, z int
	fmt.Fscan(os.Stdin, &x, &y, &z)

	var firstItem any = Point{X: x, Y: y}
	var secondItem any = z

	render(firstItem)
	render(secondItem)
}
