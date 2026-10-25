package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	var taskID int
	var taskTitle string
	fmt.Scan(&taskID, &taskTitle)

	// TODO: Создайте переменную t типа *Task через &Task{...} с именованными полями в многострочном виде (каждое поле на своей строке, с запятой после каждого поля).
	t := &Task{
		ID:    taskID,
		Title: taskTitle,
	}

	// TODO: Измените Done отдельным присваиванием после создания структуры (не в литерале): t.Done = true.
	t.Done = true

	// TODO: Выведите итоговое состояние задачи, обращаясь к полям через указатель: t.ID, t.Title, t.Done.
	fmt.Println(t.ID, t.Title, t.Done)
}
