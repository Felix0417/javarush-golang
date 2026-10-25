package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	// Множество уникальных задач: ключ — вся структура целиком.
	uniq := make(map[Task]struct{})

	for i := 0; i < n; i++ {
		var id int
		var title string
		var doneFlag int
		fmt.Fscan(in, &id, &title, &doneFlag)

		t := Task{
			ID:    id,
			Title: title,
			// TODO: Преобразуйте doneFlag (0/1) в bool и корректно заполните поле Done.
			Done: doneFlag == 1,
		}

		uniq[t] = struct{}{}
	}

	fmt.Print(len(uniq))
}
