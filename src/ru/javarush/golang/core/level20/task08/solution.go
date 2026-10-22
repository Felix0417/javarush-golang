package main

import (
	"fmt"
	"strings"
)

func normalizeTitle(raw string) string {
	/* TODO: Реализуйте нормализацию заголовка:
	   1) заменить '_' на пробелы,
	   2) затем strings.TrimSpace,
	   3) затем strings.Fields,
	   4) затем strings.Join с одиночными пробелами.
	   Важно: если после нормализации заголовок пустой — main должен вывести ошибку.
	*/
	raw = strings.ReplaceAll(raw, "_", " ")
	raw = strings.TrimSpace(raw)
	return strings.Join(strings.Fields(raw), " ")
}

func main() {
	var taskNumber int
	var doneFlag int
	var rawTitle string

	fmt.Scan(&taskNumber, &doneFlag, &rawTitle)

	title := normalizeTitle(rawTitle)
	if title == "" {
		// По условию: ровно одна строка и ничего больше
		fmt.Println("error: empty title")
		return
	}

	var b strings.Builder

	// TODO: Соберите строку результата через strings.Builder в формате:
	// <taskNumber>) [x] <title> при doneFlag == 1
	// <taskNumber>) [ ] <title> при doneFlag == 0
	//
	// Подсказка по каркасу: используйте fmt.Fprintf(&b, ...) и b.WriteString(...).

	fmt.Fprintf(&b, "%d) ", taskNumber)
	if doneFlag == 1 {
		b.WriteString("[x] ")
	} else {
		b.WriteString("[ ] ")
	}

	b.WriteString(title)

	fmt.Println(b.String())

	// TODO: Выведите второй строкой отладочное представление заголовка в формате:
	// titleDebug="<...>" используя fmt.Printf и плейсхолдер %q для title.
	fmt.Printf("titleDebug=%q\n", title)
}
