package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

// FullName возвращает имя и фамилию в "человеческом" виде: через один пробел.
func (u User) FullName() string {
	// TODO: Реализуйте форматирование полного имени в виде "<FirstName> <LastName>" с ровно одним пробелом.
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func main() {
	var firstName, lastName string
	fmt.Scan(&firstName, &lastName)

	u := User{
		FirstName: firstName,
		LastName:  lastName,
	}

	fmt.Println(u.FullName())
}
