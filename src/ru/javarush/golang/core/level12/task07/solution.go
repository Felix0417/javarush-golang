package main

import "fmt"

func main() {
	tasks := []string{"A", "B", "C", "D"}
	head := tasks[:2]

	fmt.Printf("tasks=%v len=%d cap=%d\n", tasks, len(tasks), cap(tasks))
	fmt.Printf("head=%v len=%d cap=%d\n", head, len(head), cap(head))

	// TODO: Добавьте к head элемент "X" через append и сохраните результат обратно в head.
	head = append(head, "X")

	// TODO: Выведите tasks и head после append тем же форматом (2 строки: tasks..., head...), что и до append.
	fmt.Printf("tasks=%v len=%d cap=%d\n", tasks, len(tasks), cap(tasks))
	fmt.Printf("head=%v len=%d cap=%d\n", head, len(head), cap(head))

}
