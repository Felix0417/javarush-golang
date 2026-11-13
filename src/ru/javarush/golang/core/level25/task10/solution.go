package main

/*
Вы пишете маленький модуль для списка гостей. Иногда переменная any вообще не содержит значения (это nil interface), а
иногда содержит “гостя” как Person или *Person. И есть самый хитрый случай: внутри any лежит тип *Person, но сам
указатель nil — выглядит как “что-то есть”, но гостя нет.

Создайте тип Person с полем Name string. Напишите функцию Kind(x any) string, которая различает случаи через type switch:
- если x — nil интерфейс, вернуть "nil interface";
- если внутри x лежит *Person, то при v == nil вернуть "typed nil *Person", иначе вернуть "person pointer:<имя>";
- если внутри x лежит Person (значение, не указатель), вернуть "person value:<имя>";
- для остальных типов вернуть "other:<type>" (через %T от исходного x).

В main выведите 4 строки для значений:
(a) var x any
(b) var p *Person = nil; x = p
(c) x = &Person{Name:"Ann"}
(d) x = Person{Name:"Bob"}

Ожидаемый вывод:
nil interface
typed nil *Person
person pointer:Ann
person value:Bob
*/

import "fmt"

type Person struct {
	Name string
}

func Kind(x any) string {
	switch v := x.(type) {
	case nil:
		// TODO: Верните строку для случая, когда x — nil interface.
		return "nil interface"
	case *Person:
		// TODO: Обработайте случай *Person, включая typed nil (когда v == nil).
		// TODO: Для ненулевого указателя верните строку с именем гостя.
		if v == nil {
			return "typed nil *Person"
		}
		return fmt.Sprintf("person pointer:%s", v.Name)
	case Person:
		// TODO: Обработайте случай Person-значения (не указатель) и верните строку с именем.
		return fmt.Sprintf("person value:%s", v.Name)
	default:
		// TODO: Для остальных типов верните строку вида other:<type>, где <type> получен через %T от исходного x.
		return fmt.Sprintf("other:%T", x)
	}
}

func main() {
	var x any
	fmt.Println(Kind(x))

	var p *Person = nil
	x = p
	fmt.Println(Kind(x))

	x = &Person{Name: "Ann"}
	fmt.Println(Kind(x))

	x = Person{Name: "Bob"}
	fmt.Println(Kind(x))
}
