package main

import "fmt"

func main() {
	var n int
	var text string
	fmt.Scan(&n, &text)

	// По условию: при N <= 0 печатаем пустую строку.
	if n <= 0 {
		fmt.Println()
		return
	}

	// TODO: Реализуйте обрезку по первым N Unicode-рунам, а не по байтам.
	// TODO: Нельзя преобразовывать строку в []rune.
	// TODO: Используйте for range по строке text, чтобы найти байтовый индекс начала (N+1)-й руны.
	// TODO: Если в строке меньше либо ровно N рун — выведите строку целиком без изменений.

	cut := len(text)
	var count int
	for i, _ := range text {
		if count >= n {
			cut = i
			break
		}
		count++
	}

	fmt.Println(text[:cut])
}
