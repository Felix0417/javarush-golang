package main

/*
Представьте, что обычная задача живёт «в голове», а вот когда вы сохраняете её в базе, к ней добавляется ID. Вы хотите,
чтобы в сохранённой записи поля задачи по-прежнему были доступны коротко, без цепочек .Task.Meta..., поэтому используете
embedding сразу на двух уровнях.

Программа читает из stdin четыре значения: id (целое число), затем три слова title, author, project. Создайте StoredTask,
который содержит поле ID и embedded-поле Task. Внутри Task тоже должен быть embedded-Meta.

После создания выведите одну строку: id=<ID>; title=<TITLE>; author=<AUTHOR>; project=<PROJECT>, причём title/author/project
берите именно через короткий доступ stored.Title, stored.Author, stored.Project.
*/

import "fmt"

type Meta struct {
	Title   string
	Author  string
	Project string
}

type Task struct {
	// TODO: Сделайте Meta embedded-полем (анонимным), чтобы поля Meta "поднимались" в Task.
	Meta
}

type StoredTask struct {
	ID int

	// TODO: Сделайте Task embedded-полем (анонимным), чтобы поля Task (и транзитивно Meta) "поднимались" в StoredTask.
	Task
}

func main() {
	var (
		id      int
		title   string
		author  string
		project string
	)

	fmt.Scan(&id, &title, &author, &project)

	stored := StoredTask{
		ID: id,
		Task: Task{
			Meta: Meta{
				Title:   title,
				Author:  author,
				Project: project,
			},
		},
	}

	// TODO: Выведите значения title/author/project через короткий доступ: stored.Title, stored.Author, stored.Project.
	fmt.Printf("id=%d; title=%s; author=%s; project=%s", stored.ID, stored.Title, stored.Author, stored.Project)
}
