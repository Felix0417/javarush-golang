package main

import (
	"fmt"
	"strconv"
)

func parseIntSafe(s string) (n int, ok bool) {
	// TODO: Реализуйте "мягкий" парсинг строки в int через strconv.Atoi.
	// TODO: При успехе верните (n=<число>, ok=true), при ошибке (n=0, ok=false).
	// TODO: Не используйте := при присваивании именованным результатам n и ok (только =).
	// TODO: В конце используйте return без значений.

	var err error

	// Вызов оставлен как часть каркаса, но результат пока не используется правильно.
	n, err = strconv.Atoi(s)
	if err != nil {
		n = 0
		ok = false
	} else {
		ok = true
	}
	return
}

func main() {
	var rawText string
	fmt.Scan(&rawText)

	n, ok := parseIntSafe(rawText)
	if ok {
		fmt.Printf("ok: %d", n)
	} else {
		fmt.Print("error")
	}
}
