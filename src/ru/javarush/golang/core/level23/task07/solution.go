package main

/*
Вы делаете консольный трекер задач, где статус хранится не как bool, а как отдельный тип: так проще расширять поведение
и держать логику рядом со статусом. При этом сама задача не должна «знать», как именно статус превращается в done, — она
просто делегирует это статусу.

Программа читает из stdin два слова: title и statusText. Допустимые статусы: "todo" и "done". Объявите тип Status на базе
string, задайте константы StatusTodo и StatusDone, и реализуйте метод (s Status) IsDone() bool.

Структура Task должна содержать Title и Status, а метод (t Task) Done() bool должен возвращать результат делегирования
t.Status.IsDone().

Если statusText не равен "todo" или "done", программа должна вывести "invalid status". Иначе программа должна вывести
done=<true|false>; title=<TITLE>.
*/

import (
	"errors"
	"fmt"
)

type Status string

const (
	StatusTodo Status = "todo"
	StatusDone Status = "done"
)

// Логика "задача сделана?" живёт рядом со статусом.
func (s Status) IsDone() bool {
	// TODO: Реализуйте логику определения "сделано ли" для статуса.
	return s == StatusDone
}

type Task struct {
	Title  string
	Status Status
}

// Task не знает правил done — просто делегирует статусу.
func (t Task) Done() bool {
	return t.Status.IsDone()
}

var errInvalidStatus = errors.New("invalid status")

func ParseStatus(text string) (Status, error) {
	// TODO: Реализуйте разбор текста статуса: принимайте только "todo" и "done", иначе возвращайте errInvalidStatus.
	switch Status(text) {
	case StatusTodo:
		return StatusTodo, nil
	case StatusDone:
		return StatusDone, nil
	default:
		return "", errInvalidStatus
	}
}

func main() {
	var title, statusText string
	fmt.Scan(&title, &statusText)

	status, err := ParseStatus(statusText)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	task := Task{
		Title:  title,
		Status: status,
	}

	fmt.Printf("done=%t; title=%s", task.Done(), task.Title)
}
