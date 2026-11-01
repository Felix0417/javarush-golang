package main

import "fmt"

// Task — минимальная модель задачи.
type Task struct {
	ID    int
	Title string
	Done  bool
}

// NewTask создаёт задачу с Done=false (по условию "в начале не сделана").
func NewTask(id int, title string) Task {
	return Task{
		ID:    id,
		Title: title,
	}
}

// Rename меняет название задачи (нужно менять состояние).
func (t *Task) Rename(title string) {
	// TODO: Измените поле Title у задачи на переданный title.
	t.Title = title
}

// MarkDone помечает задачу выполненной (нужно менять состояние).
func (t *Task) MarkDone() {
	// TODO: Установите поле Done в true.
	t.Done = true
}

// Label возвращает компактную строку для вставки в лог/чат (без изменения состояния).
func (t Task) Label() string {
	// TODO: Реализуйте форматирование метки:
	// TODO: - если Done == false, один формат;
	// TODO: - если Done == true, другой формат;
	// TODO: Строка должна содержать статус, ID и Title.
	status := "TODO"
	if t.Done {
		status = "DONE"
	}
	return fmt.Sprintf("[%s] #%d %s", status, t.ID, t.Title)
}

func main() {
	var taskID int
	var oldTitle, newTitle string

	fmt.Scan(&taskID, &oldTitle, &newTitle)

	task := NewTask(taskID, oldTitle)
	task.Rename(newTitle)
	task.MarkDone()

	fmt.Println(task.Label())
}
