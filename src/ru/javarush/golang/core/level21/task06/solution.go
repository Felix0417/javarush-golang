package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	var firstName, lastName string
	fmt.Scan(&firstName, &lastName)

	// TODO: Создайте переменную u типа User через именованный литерал (User{...}).
	// TODO: Важно: задайте поля НЕ в порядке объявления в структуре — сначала LastName, затем FirstName.
	u := User{
		LastName:  lastName,
		FirstName: firstName,
	}

	fmt.Printf("First=%s Last=%s", u.FirstName, u.LastName)
}
