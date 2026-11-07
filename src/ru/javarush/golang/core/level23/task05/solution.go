package main

import "fmt"

/*
Вы пишете мини-консольный прототип для команды: задача хранит заголовок отдельно, а информацию «кто создал и к какому
проекту относится» — внутри отдельной структуры метаданных. Здесь принципиально важно, что метаданные лежат не как embedded,
а как обычное именованное поле, чтобы доступ был только через явный путь.

Опишите структуру Meta с полями Author и Project типа string. Опишите структуру Task с полем Title типа string и полем Meta
типа Meta (именованное поле).

Программа читает из stdin три слова: title, author, project. На их основе создаётся значение Task, где заголовок лежит в
task.Title, а автор и проект — внутри task.Meta. Затем выведите одну строку: title=<TITLE>; author=<AUTHOR>; project=<PROJECT>.
*/

type Meta struct {
	// TODO: Добавьте поля Author и Project типа string.
	Author  string
	Project string
}

type Task struct {
	// TODO: Добавьте поле Title типа string.
	// TODO: Добавьте именованное поле Meta типа Meta (не embedded), чтобы доступ был только через task.Meta.<Field>.
	Title string
	Meta  Meta
}

func main() {
	// TODO: Прочитайте из stdin три слова (title, author, project) с помощью fmt.Scan.
	var title, author, project string
	if _, err := fmt.Scan(&title, &author, &project); err != nil {
		return
	}

	// TODO: Создайте значение task типа Task, записав:
	// TODO: - title в task.Title
	// TODO: - author в task.Meta.Author
	// TODO: - project в task.Meta.Project
	task := Task{Title: title, Meta: Meta{Author: author, Project: project}}

	// TODO: Выведите ровно одну строку в формате:
	// TODO: title=<TITLE>; author=<AUTHOR>; project=<PROJECT>
	// TODO: При выводе обращайтесь к метаданным явно: task.Meta.Author и task.Meta.Project.
	fmt.Printf("title=%s; author=%s; project=%s\n", task.Title, task.Meta.Author, task.Meta.Project)
}
