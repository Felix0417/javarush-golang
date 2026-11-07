package main

/*
Вы делаете консольный «пропуск» для учебного центра. Иногда системе нужно вывести только имя (например, на бейджике),
а иногда — полный профиль: имя и возраст (например, в отчёте). Чтобы не заставлять функции требовать лишнее, вы разделяете контракт на узкие интерфейсы и объединяете их через embedding, когда нужно больше информации.

Программа читает из stdin studentName и studentAge (строка и int) и выводит две строки.
Объявите интерфейс Namer с методом Name() string и интерфейс Ager с методом Age() int. Объявите интерфейс Profile, который
включает Namer и Ager через embedding. Создайте структуру Student с полями для имени и возраста и реализуйте на ней методы Name и Age.

Сделайте две функции: PrintName(n Namer) и PrintProfile(p Profile). Первая печатает name=<имя>. Вторая печатает name=<имя> age=<возраст>.
В main создайте Student и вызовите обе функции.
*/

import "fmt"

// Узкий контракт: нужно только имя.
type Namer interface {
	Name() string
}

// Узкий контракт: нужен только возраст.
type Ager interface {
	Age() int
}

// Сборка более широкого контракта через embedding интерфейсов.
type Profile interface {
	Namer
	Ager
}

type Student struct {
	name string
	age  int
}

func (s Student) Name() string {
	// TODO: Верните имя студента из структуры Student.
	return s.name
}

func (s Student) Age() int {
	// TODO: Верните возраст студента из структуры Student.
	return s.age
}

func PrintName(n Namer) {
	// TODO: Напечатайте только имя по контракту Namer (не используя Age) строго в формате:
	// name=<имя>
	fmt.Printf("name=%s\n", n.Name())
}

func PrintProfile(p Profile) {
	// TODO: Напечатайте имя и возраст по контракту Profile строго в формате:
	// name=<имя> age=<возраст>
	fmt.Printf("name=%s age=%d\n", p.Name(), p.Age())
}

func main() {
	var studentName string
	var studentAge int
	fmt.Scan(&studentName, &studentAge)

	student := Student{name: studentName, age: studentAge}

	PrintName(student)
	PrintProfile(student)
}
