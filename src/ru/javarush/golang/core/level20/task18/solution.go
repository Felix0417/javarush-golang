package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// TaskTitle — доменный тип: строка, которая обязана быть непустой после нормализации.
type TaskTitle string

// NewTaskTitle нормализует "сырой" заголовок.
// Инвариант: после нормализации заголовок не должен стать пустым.
func NewTaskTitle(raw string) (TaskTitle, error) {
	// TODO: Нормализуйте raw строго через strings.Fields(raw) и strings.Join(parts, " "),
	// TODO: чтобы между словами был ровно один пробел (любой пробельный мусор нужно схлопнуть).
	// TODO: После нормализации проверьте инвариант: результат не должен быть пустым, иначе верните ошибку.

	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", errors.New("empty task title")
	}
	trimmed = strings.Join(strings.Fields(trimmed), " ")

	return TaskTitle(trimmed), nil
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		fmt.Print("ERROR")
		return
	}

	// По требованиям: при N == 0 обязаны вывести ERROR.
	if n == 0 {
		fmt.Print("ERROR")
		return
	}

	tokens := make([]string, n)
	for i := 0; i < n; i++ {
		// "Попытаться прочитать N слов": если не получилось — считаем это ошибкой.
		if _, err := fmt.Fscan(in, &tokens[i]); err != nil {
			fmt.Print("ERROR")
			return
		}
	}

	// В main "сырой" заголовок собираем через три пробела.
	raw := strings.Join(tokens, "   ")

	title, err := NewTaskTitle(raw)
	if err != nil {
		fmt.Print("ERROR")
		return
	}

	fmt.Print(string(title))
}
