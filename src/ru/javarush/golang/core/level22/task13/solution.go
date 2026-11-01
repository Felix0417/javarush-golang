package main

import "fmt"

type UserCard struct {
	Name string
	Age  int
}

func (u UserCard) String() string {
	// TODO: Реализуйте строковое представление визитки, чтобы fmt.Println(card) выводил нужный формат.
	// TODO: Внутри String() нельзя печатать в stdout/stderr — метод должен только вернуть строку.
	return fmt.Sprintf("User %q, age=%d", u.Name, u.Age)
}

func main() {
	var userName string
	var userAge int
	fmt.Scan(&userName, &userAge)

	card := UserCard{Name: userName, Age: userAge}
	fmt.Println(card)
}
