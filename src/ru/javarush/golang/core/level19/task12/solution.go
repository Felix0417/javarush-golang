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

func NewTask(id int, title string) *Task {
	// По условию: создаём именно через new(Task)
	t := new(Task)

	// TODO: заполните поля t.ID и t.Title значениями id и title
	// Done должен остаться в zero value (false)
	t.ID = id
	t.Title = title

	return t
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	// TODO: создайте слайс tasks типа []*Task через make([]*Task, 0, n)
	tasks := make([]*Task, 0, n)

	byID := make(map[int]*Task)

	for i := 0; i < n; i++ {
		var id int
		var title string
		fmt.Fscan(in, &id, &title)

		t := NewTask(id, title)
		tasks = append(tasks, t)

		// TODO: в карту byID нужно положить тот же указатель, что и в tasks (t), а не создавать новый объект
		byID[id] = t
	}

	var doneID int
	fmt.Fscan(in, &doneID)

	if _, ok := byID[doneID]; ok {
		// TODO: убедитесь, что отметка Done=true делается через указатель из byID и отражается в tasks
		byID[doneID].Done = true
	}

	for _, t := range tasks {
		fmt.Fprintln(out, t.ID, t.Title, t.Done)
	}
}
