package main

import (
	"errors"
	"fmt"
	"strings"
)

// updateAt переименовывает задачу по индексу.
func updateAt(tasks []string, idx int, title string) ([]string, error) {
	// TODO: Реализуйте проверку корректности индекса idx до любого обращения к tasks[idx], чтобы не было panic.
	// TODO: Если idx некорректный — верните исходный tasks без изменений и ошибку с текстом "index out of range".
	// TODO: Если idx корректный — замените элемент по индексу на title и верните обновлённый слайс и nil.
	if idx < 0 || idx >= len(tasks) {
		return tasks, errors.New("index out of range")
	}

	tasks[idx] = title
	return tasks, nil
}

func main() {
	tasks := []string{"read", "code", "sleep"}

	var idx int
	var title string
	fmt.Scan(&idx, &title)

	updated, err := updateAt(tasks, idx, title)
	if err != nil {
		fmt.Print("error: index out of range")
		return
	}

	fmt.Print("updated: " + strings.Join(updated, ","))
}
