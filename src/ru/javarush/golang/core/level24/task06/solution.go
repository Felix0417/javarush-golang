package main

/*
Представьте, что у вас есть “визитка пользователя” — тип User. Иногда ссылка на пользователя может быть nil (пользователь
ещё не создан), но вы всё равно передаёте это значение как интерфейс fmt.Stringer, чтобы печать была единой. И тут
начинается магия Go: указатель nil может превратиться в интерфейс “не nil”.

Создайте структуру User с полем Name. Реализуйте метод String() string с receiver’ом *User. Метод должен быть
nil‑устойчивым: если receiver равен nil, возвращайте строку "<nil User>", иначе возвращайте строку с именем (любой
понятный формат).

В main объявите var u *User = nil, затем присвойте u в переменную интерфейсного типа fmt.Stringer. Выведите
сравнения (u == nil) и (s == nil), а затем выведите fmt.Printf("%T %v\n", s, s), чтобы увидеть и тип, и значение.
Создайте структуру User с полем Name.
Реализуйте метод String() string с receiver’ом *User.
В main объявите var u *User = nil, затем присвойте u в переменную интерфейсного типа fmt.Stringer.
Выведите сравнения (u == nil) и (s == nil), а затем выведите fmt.Printf("%T %v\n", s, s).
*/

import "fmt"

type User struct {
	Name string
}

func (u *User) String() string {
	// TODO: Реализуйте nil-устойчивое строковое представление пользователя.
	// TODO: Если receiver == nil, верните строку "<nil User>" (без panic).
	// TODO: Иначе верните строку, содержащую имя пользователя в любом понятном формате.
	if u == nil {
		return "<nil User>"
	}
	return u.Name
}

func main() {
	var u *User = nil

	var s fmt.Stringer = u

	// TODO: Выведите результаты сравнений (u == nil) и (s == nil) (две отдельные строки).
	// TODO: Затем выведите тип и значение интерфейса через fmt.Printf("%T %v\n", s, s).
	fmt.Println(u == nil)
	fmt.Println(s == nil)
	fmt.Printf("%T %v\n", s, s)
}
