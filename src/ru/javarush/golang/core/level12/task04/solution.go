package main

import "fmt"

func main() {
	// Заранее заданный список задач (менять нельзя).
	tasks := []string{"learn Go", "write code", "drink water", "sleep", "repeat"}

	var from, to int
	fmt.Scan(&from, &to)

	// Переводим "человеческие" границы (1-based, включительно) в границы среза.
	a := from - 1
	b := to

	// TODO: Реализуйте проверку корректности диапазона ДО нарезки среза, чтобы исключить panic.
	// TODO: Если диапазон некорректный — выведите ровно одну строку: [] и завершите программу.
	// TODO: Если диапазон корректный — выберите задачи через срез и выведите их построчно в формате "<номер>: <текст>".
	if from < 1 || to < 1 || from > to || to > len(tasks) {
		fmt.Println([]string{})
		return
	}

	s := tasks[a:b]

	for i, v := range s {
		fmt.Printf("%d: %s\n", i+1, v)
	}
}
