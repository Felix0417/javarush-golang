package main

/*
Вы делаете рендер для консольного магазина. У некоторых типов есть “особый красивый вывод” через fmt.Stringer, но для Price
вы хотите прям совсем особый приоритет: даже если он тоже Stringer, его нужно обработать своей веткой раньше, чтобы
формат был строго контролируемый.

Создайте тип Price с полем Cents int и методом String() string, который возвращает цену в формате "$D.CC" (например,
при 199 центах строка должна быть "$1.99").

Также создайте тип Note на базе string и реализуйте для него String() string, возвращая note(<text>).

Напишите функцию Render(x any) string, которая выбирает логику через type switch в таком порядке:
- case Price: и вернуть price: + v.String()
- case fmt.Stringer: и вернуть stringer: + v.String()
- case string: и вернуть text: + строку
- case nil: вернуть nil
- иначе вернуть unknown:<type> через %T от исходного x

В main выведите построчно результат для значений: Price{Cents:199}, Note("hi"), "raw", 10.
Ожидаемый вывод:
price:$1.99
stringer:note(hi)
text:raw
unknown:int
*/

import (
	"fmt"
)

type Price struct {
	Cents int
}

func (p Price) String() string {
	// TODO: Реализуйте форматирование цены из p.Cents в строку вида "$D.CC" (центы всегда двумя цифрами).
	dollars := p.Cents / 100
	cents := p.Cents % 100
	return fmt.Sprintf("$%d.%02d", dollars, cents)
}

type Note string

func (n Note) String() string {
	// TODO: Реализуйте человекочитаемый формат заметки: "note(<text>)".
	return fmt.Sprintf("note(%s)", string(n))
}

func Render(x any) string {
	// TODO: Реализуйте выбор логики через type switch (v := x.(type)) с нужным приоритетом веток.
	// TODO: Учтите, что Price должен обрабатываться раньше fmt.Stringer.
	switch v := x.(type) {
	case Price:
		// TODO: Верните строку для Price.
		return fmt.Sprintf("price:%v", v.String())
	case fmt.Stringer:
		// TODO: Верните строку для fmt.Stringer.
		return fmt.Sprintf("stringer:%v", v.String())
	case string:
		// TODO: Верните строку для обычного текста.
		return fmt.Sprintf("text:%s", v)
	case nil:
		// TODO: Верните строку для nil.
		return "nil"
	default:
		// TODO: Верните строку для неизвестного типа, используя %T от исходного x.
		return fmt.Sprintf("unknown:%T", x)
	}
}

func main() {
	fmt.Println(Render(Price{Cents: 199}))
	fmt.Println(Render(Note("hi")))
	fmt.Println(Render("raw"))
	fmt.Println(Render(10))
}
