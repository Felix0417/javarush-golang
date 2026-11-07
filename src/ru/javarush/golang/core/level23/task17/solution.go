package main

/*
Вы делаете крошечный прототип для управления домашним проектом: каждому новому делу нужен уникальный номер, и вы хотите,
чтобы проект сам умел «выдавать следующий ID» одной командой, без лишних уточнений. Для этого вы прячете генератор внутрь
проекта как embedded-часть и используете promoted-метод.

Напишите программу с типом IDGen, который хранит поле next (int) и имеет метод Next() int (pointer receiver). Метод
увеличивает next на 1 и возвращает новое значение.

Создайте тип Project с полем Name (string) и встроенным (embedded) IDGen. В функции main создайте проект с именем "Home"
и вызовите Next() три раза через переменную проекта (именно как p.Next()).
Программа должна выводить 4 строки:
Home
1
2
3
*/

import "fmt"

// IDGen хранит состояние генератора ID.
type IDGen struct {
	next int
}

// Next возвращает следующий ID и обновляет состояние генератора.
func (g *IDGen) Next() int {
	// TODO: Реализуйте генерацию следующего ID: увеличить g.next на 1 и вернуть новое значение (метод должен менять состояние).
	g.next++
	return g.next
}

// Project встраивает IDGen, чтобы Next() был promoted-методом.
type Project struct {
	Name string
	IDGen
}

func main() {
	p := Project{Name: "Home"}

	fmt.Println(p.Name)
	fmt.Println(p.Next())
	fmt.Println(p.Next())
	fmt.Println(p.Next())
}
