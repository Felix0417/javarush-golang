package main

import (
	"fmt"
	"strings"
)

func FindTaskPtrs(tasks []string, query string) []*string {
	var res []*string

	// query будет нужен в условии поиска.

	for i := range tasks {
		// TODO: Реализуйте проверку, что tasks[i] содержит подстроку query через strings.Contains.
		// TODO: Если условие истинно — добавьте в res указатель на элемент исходного слайса: &tasks[i].
		if strings.Contains(tasks[i], query) {
			res = append(res, &tasks[i])
		}
	}

	// Подсказка компилятору: строковые функции должны использоваться в решении.

	return res
}

func main() {
	tasks := []string{"buy_milk", "learn_go", "sleep", "go_gym"}

	var query string
	fmt.Scan(&query)

	ptrs := FindTaskPtrs(tasks, query)
	for _, p := range ptrs {
		// TODO: Преобразуйте строку по указателю p в верхний регистр через strings.ToUpper и запишите обратно.
		*p = strings.ToUpper(*p)
	}

	// Подсказка компилятору: строковые функции должны использоваться в решении.

	fmt.Println(tasks)
}
