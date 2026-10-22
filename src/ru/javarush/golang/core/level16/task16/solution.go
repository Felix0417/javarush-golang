package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Про стиль ошибок в Go (строчные, без точки) см.

func main() {
	// Задачи храним отдельно от статусов, без структур (требование).
	tasks := map[int]string{
		1: "купить молоко",
		2: "помыть посуду",
		3: "прочитать книгу",
	}
	done := map[int]bool{
		1: false,
		2: false,
		3: false,
	}

	var cmd, idStr string
	fmt.Fscan(os.Stdin, &cmd, &idStr)

	if cmd != "done" {
		fmt.Println(fmt.Errorf("unknown command %q", cmd))
		return
	}

	if err := handleDone(tasks, done, idStr); err != nil {
		fmt.Println(err)
	}
}

func handleDone(tasks map[int]string, done map[int]bool, idStr string) error {
	id, err := parseID(idStr)
	if err != nil {
		return err
	}

	// TODO: Проверьте, что задача с таким id существует в tasks; если нет — верните ошибку в нужном формате.
	// TODO: Отметьте задачу выполненной в done и выведите "ok" (одной строкой).
	_, ok := tasks[id]
	if !ok {
		return errors.New("not implemented")
	}

	done[id] = true
	fmt.Println("ok")
	return nil
}

func parseID(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		// TODO: Верните ошибку парсинга в нужном формате (со значением строки через %q), без panic.
		return 0, fmt.Errorf("invalid id")
	}
	return id, nil
}
