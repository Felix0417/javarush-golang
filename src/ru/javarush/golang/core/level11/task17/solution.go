package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	// TODO: Инициализируйте tasks как пустой по длине слайс с заранее зарезервированной ёмкостью под n.
	tasks := make([]string, 0, n)

	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(in, &word)

		// TODO: Добавьте считанное слово word в слайс tasks, увеличивая длину только через append.
		tasks = append(tasks, word)
	}

	fmt.Printf("len=%d cap=%d\n", len(tasks), cap(tasks))

	if n == 0 {
		fmt.Print("\n")
		return
	}

	for i, task := range tasks {
		// TODO: Выведите задачи построчно с нумерацией с 1 в формате "<номер>) <слово>".
		fmt.Printf("%d) %s\n", i+1, task)
	}
}
