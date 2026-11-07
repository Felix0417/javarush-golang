package main

import "fmt"

type Meta struct {
	ID int
}

type Task struct {
	Title string
	// TODO: Сделайте поле Meta embedded (анонимным), чтобы ID "поднялся" в Task.
	Meta
}

func main() {
	task := Task{
		Title: "LearnEmbedding",
		// TODO: Задайте ID = 42. Promoted-поле ID ключом литерала Task не задаётся — только через embedded-поле Meta.
		Meta: Meta{
			ID: 42,
		},
	}

	// TODO: Выведите ID через короткий доступ task.ID (после embedding) — это должна быть первая строка.
	fmt.Println(task.ID)

	// Вторая строка должна выводить то же значение, но через явный путь task.Meta.ID.
	fmt.Println(task.Meta.ID)
}
