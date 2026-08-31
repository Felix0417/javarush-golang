package main

import "fmt"

func main() {
	var s, repl string
	var i int

	// Читаем ровно три значения (s, i, repl).
	fmt.Scan(&s, &i, &repl)

	// Важно: сначала проверяем границы и repl, и только потом делаем срезы/индексацию.
	if i < 0 || i >= len(s) || len(repl) == 0 {
		fmt.Print("OUT_OF_RANGE")
		return
	}

	// TODO: Реализуйте замену ровно одного байта строки s в позиции i на первый байт строки repl.
	// TODO: Соберите результат через срезы и конкатенацию трех частей (без преобразования строк в []byte).
	// TODO: Нельзя использовать пакеты strings/bytes/strconv — допускается только fmt.

	// Временная заглушка: выводим исходную строку без изменений.
	fmt.Print(s[:i] + string(repl[0]) + s[i+1:])
}
