package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	// TODO: Посчитайте количество рун (Unicode-символов) через range по строке.
	// Важно: нельзя использовать utf8.RuneCountInString и нельзя импортировать другие пакеты.
	runes := 0
	for range s {
		// TODO: Увеличьте счётчик рун внутри цикла.
		runes++
	}

	fmt.Printf("bytes=%d runes=%d\n", len(s), runes)

	// TODO: Выведите построчный отчёт по каждой руне через `for i, r := range s`.
	// i должен быть байтовым индексом начала руны (значение i из range),
	// k — отдельным счётчиком рун, начиная с 1.
	k := 1
	for i, r := range s {
		// TODO: Напечатайте строку в точном формате:
		// k=<номер_руны_с_1> i=<байтовый_индекс> code=<код_руны> char=<символ>
		// code и char должны быть получены из текущей руны r.
		fmt.Printf("k=%d i=%d code=%d char=%c\n", k, i, r, r)

		k++
	}
}
