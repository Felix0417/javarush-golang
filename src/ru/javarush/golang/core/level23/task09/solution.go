package main

/*
Вы пишете простого чат-бота для консоли: он должен уметь и поздороваться, и попрощаться. В команде договорились описывать
возможности через маленькие интерфейсы, а потом собирать общий контракт через embedding, чтобы типы легко комбинировались
и тестировались.

Программа читает из stdin одно слово — userName — и печатает две строки приветствия.
В коде объявите два интерфейса: первый описывает метод Hello(name string) string, второй — метод Bye(name string) string.
Затем объявите третий интерфейс Greeter, который собирает общий контракт через embedding этих двух интерфейсов.

Создайте тип SimpleGreeter, который реализует оба метода. В main создайте переменную интерфейсного типа Greeter, присвойте
ей SimpleGreeter и выведите результат вызова Hello и Bye (каждый — с новой строки).
*/

import "fmt"

type Helloer interface {
	Hello(name string) string
}

type Byer interface {
	Bye(name string) string
}

// Greeter собирается через embedding маленьких интерфейсов, а не через перечисление методов.
type Greeter interface {
	Helloer
	Byer
}

type SimpleGreeter struct{}

func (SimpleGreeter) Hello(name string) string {
	// TODO: Реализуйте формирование строки приветствия с использованием имени.
	return "Hello, " + name
}

func (SimpleGreeter) Bye(name string) string {
	// TODO: Реализуйте формирование строки прощания с использованием имени.
	return "Bye " + name
}

func main() {
	var userName string
	fmt.Scan(&userName)

	// Важно: используем переменную интерфейсного типа для вызова методов.
	var g Greeter = SimpleGreeter{}

	fmt.Println(g.Hello(userName))
	fmt.Println(g.Bye(userName))
}
