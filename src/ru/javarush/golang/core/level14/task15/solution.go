package main

import (
	"fmt"
	"strings"
)

func main() {
	// Исходная "грязная" строка тегов (по условию)
	rawTags := "  go,  strings,,  utf-8  ,  "

	// 1) Обрезаем пробелы по краям всей строки
	clean := strings.TrimSpace(rawTags)

	// 2) Разбиваем по запятой
	parts := strings.Split(clean, ",")

	// 3) Собираем теги (пока упрощённо)
	var tags []string
	for _, p := range parts {
		// TODO: Обрежьте пробелы у каждого тега через strings.TrimSpace (нельзя оставлять лишние пробелы внутри списка).
		// TODO: Отфильтруйте пустые теги (пустая строка после обрезки пробелов не должна попадать в итоговый список).
		p = strings.TrimSpace(p)

		// Временная упрощённая фильтрация: убирает только строго пустые фрагменты,
		// но не учитывает фрагменты из одних пробелов и не чистит пробелы вокруг тегов.
		if p == "" {
			continue
		}

		tags = append(tags, p)
	}

	// 4) Склеиваем теги через ", "
	list := strings.Join(tags, ", ")

	// 5) Собираем финальную строку через strings.Builder (без конкатенации)
	var b strings.Builder
	b.WriteString("Tags: ")
	b.WriteString(list)

	// Ровно одна строка вывода
	fmt.Println(b.String())
}
