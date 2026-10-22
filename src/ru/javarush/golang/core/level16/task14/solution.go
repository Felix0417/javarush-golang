package main

import (
	"fmt"
	"strconv"
)

func main() {
	var idText string
	fmt.Scan(&idText) // по условию: один токен без пробелов

	id, err := parseID(idText)
	if err != nil {
		fmt.Println(err) // ровно одна строка ошибки
		return
	}

	fmt.Printf("id=%d\n", id) // ровно одна строка успеха
}

func parseID(s string) (int, error) {
	id, err := strconv.Atoi(s) // обязательно strconv.Atoi
	if err != nil {
		// TODO: Сформируйте ошибку через fmt.Errorf: добавьте исходный ввод (s) в кавычках через %q и сохраните текст исходной ошибки парсинга.
		return 0, fmt.Errorf("parse id %q: %v", s, err)
	}
	return id, nil
}
