package main

import "fmt"

func main() {
	// Справочник задан заранее (не вводится пользователем).
	ages := map[string]int{
		// TODO: добавьте несколько пар (имя → возраст); обязательно добавьте хотя бы одну запись с возрастом 0
		"baby":  0,
		"child": 2,
		"teen":  13,
		"adult": 19,
	}

	var personName string
	fmt.Scan(&personName)

	// TODO: найдите возраст по имени в map через ok-идиому, чтобы различать:
	// TODO: 1) ключа нет (ok == false)
	// TODO: 2) ключ есть, но возраст может быть 0 (ok == true, age == 0)
	// TODO: в зависимости от ok выведите либо "age: <число>", либо "no such person" (точный формат)
	if v, ok := ages[personName]; ok {
		fmt.Printf("age: %d\n", v)
	} else {
		fmt.Println("no such person")
	}
}
