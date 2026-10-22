package main

import (
	"fmt"
	"os"
	"strconv"
)

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func sumTwo(s1, s2 string) (sum int, err error) {
	var token string
	// TODO: Реализовать перехват паники через defer + recover, чтобы panic из mustParseInt не доходила до main.
	// TODO: При панике вернуть error != nil (sum может быть любым, но обычно 0).
	// TODO: Сформировать содержательную ошибку: текст должен содержать проблемный токен в кавычках и значение паники.
	defer func() {
		if r := recover(); r != nil {
			sum = 0
			err = fmt.Errorf("parse int \"%s\": %v", token, r)
		}
	}()

	token = s1
	a := mustParseInt(s1)

	token = s2
	b := mustParseInt(s2)

	sum = a + b
	return sum, nil
}

func main() {
	var s1, s2 string
	_, _ = fmt.Fscan(os.Stdin, &s1, &s2)

	sum, err := sumTwo(s1, s2)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("sum: %d\n", sum)
}
