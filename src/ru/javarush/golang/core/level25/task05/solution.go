package main

/*
Вы делаете консольный “сканер токенов” для внутренней админки: оператор вводит один короткий токен без пробелов. Иногда
это число (например, ID), а иногда — слово (например, код команды). Вам нужно сохранить токен в any, но так, чтобы он
стал int, если это реально число, иначе остался строкой.

Программа должна прочитать из stdin один токен rawToken. Затем попробуйте преобразовать его в int через strconv.Atoi.
Если преобразование удалось, сохраните значение в переменную token типа any как int. Если не удалось — сохраните в token
как string.

После этого программа должна вызвать formatToken(token any), которая с помощью безопасного type assertion вида v, ok := x.(T)
формирует строку результата: для int — квадрат числа, для string — длину строки в байтах.

Итоговую строку нужно напечатать в stdout.
*/

import (
	"fmt"
	"strconv"
)

func formatToken(token any) string {
	// Безопасно проверяем: лежит ли внутри int
	if n, ok := token.(int); ok {
		// TODO: если внутри int, вернуть квадрат числа (n*n) как строку
		sqrt := n * n
		return strconv.Itoa(sqrt)
	}

	// Безопасно проверяем: лежит ли внутри string
	if s, ok := token.(string); ok {
		// TODO: если внутри string, вернуть длину строки в байтах (len(s)) как строку
		return strconv.Itoa(len(s))
	}

	// На случай неожиданного типа
	return "unknown"
}

func main() {
	var rawToken string
	fmt.Scan(&rawToken)

	n, err := strconv.Atoi(rawToken)

	var token any

	// TODO: сохранить в token значение типа int, если strconv.Atoi успешно распарсил число
	if err == nil {
		token = n
	} else {
		token = rawToken
	}

	fmt.Println(formatToken(token))
}
