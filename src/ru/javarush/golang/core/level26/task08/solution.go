package main

/*
Вы делаете мини-отчёт по задачам для простого консольного трекера. Внутри программы задачи живут как структуры, а
пользователю нужен аккуратный текстовый список. Важно: печать должна быть переиспользуемой — один раз пишем рендер, и
потом можем отправлять его куда угодно (в stdout, в файл, в буфер).

Опишите тип Task с полями ID int, Title string, Done bool. Реализуйте функцию WriteTasks(w io.Writer, tasks []Task) error,
которая печатает каждую задачу отдельной строкой в формате:
#<id> [DONE] <title> или #<id> [TODO] <title>,
и всегда добавляет перевод строки в конце.

Затем реализуйте FormatTasks(tasks []Task) (string, error): она должна собрать весь отчёт одной строкой через
strings.Builder, но при этом переиспользовать WriteTasks, передав в неё &builder.

В main() прочитайте число tasksCount, затем tasksCount задач: id (int), done (0 или 1), title (одно слово без пробелов).
После этого выведите результат FormatTasks в stdout.
*/

import (
	"fmt"
	"io"
	"strings"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

func WriteTasks(w io.Writer, tasks []Task) error {
	// TODO: Реализуйте вывод каждой задачи в точном формате:
	// TODO: "#<id> [DONE] <title>\n" если Done=true и "#<id> [TODO] <title>\n" если Done=false.
	// TODO: Важно: выводить нужно в переданный io.Writer и возвращать ошибку записи (если она возникла).

	var taskStatus string

	for _, t := range tasks {
		if t.Done {
			taskStatus = "DONE"
		} else {
			taskStatus = "TODO"
		}

		if _, err := fmt.Fprintf(w, "#%d [%s] %s\n", t.ID, taskStatus, t.Title); err != nil {
			return err
		}
	}

	return nil
}

func FormatTasks(tasks []Task) (string, error) {
	// TODO: Соберите весь отчёт в одну строку через strings.Builder (используя нулевое значение).
	// TODO: Переиспользуйте WriteTasks, передав в неё &builder.
	// TODO: Если WriteTasks вернула ошибку — верните её из FormatTasks.

	var b strings.Builder

	if err := WriteTasks(&b, tasks); err != nil {
		return "", err
	}

	return b.String(), nil
}

func main() {
	var tasksCount int
	fmt.Scan(&tasksCount)

	tasks := make([]Task, 0, tasksCount)
	for i := 0; i < tasksCount; i++ {
		var id int
		var doneInt int
		var title string
		fmt.Scan(&id, &doneInt, &title)

		tasks = append(tasks, Task{
			ID:    id,
			Title: title,
			Done:  doneInt == 1,
		})
	}

	out, err := FormatTasks(tasks)
	if err != nil {
		return
	}
	fmt.Print(out)
}
