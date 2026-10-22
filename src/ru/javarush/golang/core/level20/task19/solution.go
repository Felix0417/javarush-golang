package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Status int

const (
	// Важно: значения по условию (iota даёт 0,1,2)
	StatusUnknown Status = iota
	StatusTodo
	StatusDone
)

func ParseStatus(s string) (Status, error) {
	// TODO: Реализуйте нормализацию строки (strings.TrimSpace + strings.ToLower) и строгое распознавание только "todo"/"done".
	// Текущая реализация намеренно неполная: она не обрабатывает разные регистры и другие варианты "грязного" ввода.

	s = strings.ToLower(strings.TrimSpace(s))

	switch s {
	case "todo":
		return StatusTodo, nil
	case "done":
		return StatusDone, nil
	default:
		return StatusUnknown, errors.New("invalid status")
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var m int
	fmt.Fscan(in, &m)

	var todo, done, invalid int
	for i := 0; i < m; i++ {
		var raw string
		fmt.Fscan(in, &raw)

		st, err := ParseStatus(raw)
		if err != nil {
			invalid++
			continue // не останавливаемся на ошибке
		}

		switch st {
		case StatusTodo:
			todo++
		case StatusDone:
			done++
		}
	}

	fmt.Fprintf(out, "todo=%d done=%d invalid=%d", todo, done, invalid)
}
