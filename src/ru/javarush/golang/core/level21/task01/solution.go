package main

import "fmt"

type Profile struct {
	ID      int
	Nick    string
	Premium bool
}

func main() {
	// Профиль только что создан: все поля должны быть в zero value.
	var p Profile

	// TODO: Выведите в первой строке тип переменной p через fmt.Printf с плейсхолдером %T (и переводом строки).
	fmt.Printf("TYPE_PLACEHOLDER=%T\n", p)

	// TODO: Выведите во второй строке значения полей p.ID, p.Nick (через %q) и p.Premium,
	// TODO: обращаясь к полям только через селектор (p.ID, p.Nick, p.Premium). Должно получиться ровно две строки вывода.
	fmt.Printf("VALUES_PLACEHOLDER=%d %q %t\n", p.ID, p.Nick, p.Premium)
}
