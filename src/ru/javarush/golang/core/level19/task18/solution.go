package main

import "fmt"

func main() {
	tasks := []string{"buy_milk", "learn_go", "sleep"}

	var currentTask string
	var ptrs []*string

	for i := range tasks {
		currentTask = tasks[i]

		// TODO: Сделайте так, чтобы в ptrs добавлялся адрес переменной currentTask (то есть &currentTask),
		// TODO: а не адрес элемента tasks. Переменная currentTask должна быть одна и переиспользоваться в цикле.
		ptrs = append(ptrs, &currentTask)
	}

	fmt.Printf("%s %s %s", *ptrs[0], *ptrs[1], *ptrs[2])
}
