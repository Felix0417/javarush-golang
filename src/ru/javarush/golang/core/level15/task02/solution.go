package main

import "fmt"

func main() {
	var contactName, firstPhone, secondPhone string
	fmt.Scan(&contactName, &firstPhone, &secondPhone)

	// Пустая телефонная книга на один контакт.
	phonebook := make(map[string]string)

	// TODO: Запишите firstPhone по ключу contactName в phonebook.
	phonebook[contactName] = firstPhone
	// TODO: Перезапишите значение по ключу contactName на secondPhone.
	phonebook[contactName] = secondPhone

	fmt.Println(phonebook[contactName])
	fmt.Println(len(phonebook))

	// TODO: Удалите контакт contactName из phonebook через delete(phonebook, contactName).
	delete(phonebook, contactName)

	fmt.Println(len(phonebook))
}
