package main

/*
Вы пишете маленькую “кнопку приветствия” для приложения: есть функция, которая печатает приветствие, и ей совершенно не
важно, кто именно генерирует текст — лишь бы умел это делать. Поэтому интерфейс должен жить рядом с функцией‑потребителем
и содержать ровно тот минимум, который нужен прямо здесь и сейчас.

В программе должна быть функция PrintGreeting(g Greeter, name string). Ей нужен только один метод — получить строку
приветствия, поэтому интерфейс Greeter должен содержать ровно один метод Greet(name string) string и быть объявлен в том
же файле рядом с PrintGreeting.

Создайте конкретный тип SimpleGreeter, который реализует этот метод и возвращает строку вида: "Hello, <name>!". В main
создайте SimpleGreeter, вызовите PrintGreeting с именем "Go" и выведите результат в stdout.
*/

import "fmt"

// Greeter объявляем рядом с потребителем: нужен ровно один метод.
type Greeter interface {
	Greet(name string) string
}

// PrintGreeting зависит только от интерфейса, а не от конкретной реализации.
func PrintGreeting(g Greeter, name string) string {
	return g.Greet(name)
}

type SimpleGreeter struct{}

func (SimpleGreeter) Greet(name string) string {
	// TODO: Реализуйте формирование строки приветствия согласно условию задачи.
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	g := SimpleGreeter{}
	fmt.Println(PrintGreeting(g, "Go"))
}
