package main

import "fmt"

// Status — пользовательский тип вместо "голых чисел".
type Status int

const (
	StatusNew        Status = 0
	StatusInProgress Status = 1
	StatusDone       Status = 2
)

func main() {
	var taskID, statusCode int
	fmt.Scan(&taskID, &statusCode)

	// Явная конверсия в Status, но исходный statusCode (int) держим отдельно для unknown-ветки.
	status := Status(statusCode)

	var text string

	// TODO: Реализуйте выбор текста статуса через if / else if / else:
	// TODO: сравнивайте status только с StatusNew / StatusInProgress / StatusDone (не с числами),
	// TODO: а для неизвестного кода формируйте "unknown status (<statusCode>)".
	if status == StatusNew {
		// TODO: Установите корректный текст для статуса "new".
		text = "new"
	} else if status == StatusInProgress {
		// TODO: Установите корректный текст для статуса "in progress".
		text = "in progress"
	} else if status == StatusDone {
		// TODO: Установите корректный текст для статуса "done".
		text = "done"
	} else {
		text = fmt.Sprintf("unknown status (%d)", statusCode)
	}

	fmt.Printf("task #%d: %s\n", taskID, text)
}
