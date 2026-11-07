package main

import "fmt"

type Meta struct {
	// TODO: Добавьте поле ID типа int.
	ID int
}

type Task struct {
	Title string
	// TODO: Встройте указатель на Meta (embedded поле *Meta), чтобы были доступны task.Meta и promoted-доступ task.ID.
	*Meta
}

func main() {
	var title string
	var id int
	fmt.Scan(&title, &id)

	task := Task{Title: title}

	// TODO: Если id != -1, создайте Meta в памяти и присвойте её в task.
	// TODO: Если id == -1, оставьте метаданные nil.
	if id != -1 {
		task.Meta = &Meta{ID: id}
	}

	fmt.Printf("title=%s\n", task.Title)

	// TODO: Выведите либо id=<значение> (используя task.ID), либо id=none, не допуская паники при nil-метаданных.
	if task.Meta == nil {
		fmt.Println("id=none")
	} else {
		fmt.Printf("id=%d\n", task.ID)
	}
}
