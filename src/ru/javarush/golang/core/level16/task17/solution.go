package main

import "fmt"

// IsValidID проверяет корректность id.
// Некорректный id — это нормальный исход, поэтому возвращаем bool, а не error.
func IsValidID(id int) bool {
	// TODO: Реализуйте проверку корректности id по условию задачи.
	return id > 0
}

func main() {
	var id int
	fmt.Scan(&id)

	if IsValidID(id) {
		fmt.Print("VALID")
	} else {
		fmt.Print("INVALID")
	}
}
