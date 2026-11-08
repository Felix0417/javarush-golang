package main

/*
Вы пишете маленький модуль "ввод положительного числа" для консольного приложения. Пользователь вводит строку, а вы
пытаетесь превратить её в целое число. Если строка не число — нужно объяснить проблему. Если число есть, но оно не
положительное — тоже нужна понятная ошибка. И всё это — через ваш собственный тип ошибки, чтобы сообщение было строгим
и предсказуемым.

Программа читает из stdin одну строку s. Вы создаёте тип ошибки PositiveIntError (структура) с полями Input string и
Problem string и реализуете у неё метод Error() string так, чтобы получалось: invalid number <input>: <problem>. Затем
пишете функцию ParsePositiveInt(s string) (int, error):
- если strconv.Atoi не смог распарсить число, возвращаете PositiveIntError с Problem="not an int";
- если число распарсилось, но оно <= 0, возвращаете PositiveIntError с Problem="must be > 0";
- иначе возвращаете число и nil.

В main вы печатаете либо число, либо текст ошибки.
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// PositiveIntError — typed error: хранит исходный ввод и причину.
type PositiveIntError struct {
	Input   string
	Problem string
}

// Error реализует интерфейс error.
func (e PositiveIntError) Error() string {
	// TODO: Сформируйте сообщение об ошибке строго по формату из условия задачи.
	return fmt.Sprintf("invalid number %s: %s", e.Input, e.Problem)
}

func ParsePositiveInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		// TODO: Верните typed error PositiveIntError с заполненными полями (вместо "сырой" ошибки парсинга).
		return 0, PositiveIntError{Input: s, Problem: "not an int"}
	}

	if n <= 0 {
		// TODO: Верните typed error PositiveIntError с заполненными полями для случая неположительного числа.
		return 0, PositiveIntError{Input: s, Problem: "must be > 0"}
	}

	return n, nil
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan() // по условию читаем ровно одну строку
	s := sc.Text()

	n, err := ParsePositiveInt(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)
}
