package main

import (
	"fmt"
)

// GetTitleByID — обычное чтение: "не найдено" = нормальный исход (T, bool).
func GetTitleByID(titles map[int]string, id int) (string, bool) {
	// TODO: Реализуйте чтение заголовка из карты titles по ключу id и верните (title, ok).
	title, ok := titles[id]
	return title, ok
}

// MarkDone — команда: "не найдено" = ошибка сценария (error).
func MarkDone(done map[int]bool, titles map[int]string, id int) error {
	// TODO: Проверьте, существует ли задача с таким id в titles. Если нет — верните ошибку с понятным текстом.
	// TODO: Если задача существует — установите done[id] = true и верните nil.
	if _, ok := titles[id]; !ok {
		return fmt.Errorf("no title for id %d", id)
	}
	done[id] = true
	return nil
}

func main() {
	// In-memory "база" задач: минимум 3 задачи с id 1,2,3.
	titles := map[int]string{
		1: "Купить молоко",
		2: "Сделать домашку",
		3: "Прочитать главу про map",
	}

	// Отметки выполнения; на старте можно пустую.
	done := make(map[int]bool)

	var cmd string
	var id int
	fmt.Scan(&cmd, &id)

	switch cmd {
	case "get":
		title, ok := GetTitleByID(titles, id)
		if !ok {
			fmt.Println("NOT_FOUND")
			return
		}
		fmt.Println(title)

	case "done":
		if err := MarkDone(done, titles, id); err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("OK")
	}
}
