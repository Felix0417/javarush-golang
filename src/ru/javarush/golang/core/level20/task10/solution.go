package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var ageText string
	fmt.Scan(&ageText)

	age, err := ParseAge(ageText)
	if err != nil {
		// По условию: печатаем только текст ошибки и перевод строки.
		fmt.Println(err.Error())
		return
	}

	fmt.Println(age)
}

func ParseAge(s string) (int, error) {
	// TODO: Реализуйте аккуратный разбор возраста по требованиям:
	// TODO: 1) нормализуйте строку (обрежьте пробелы по краям),
	// TODO: 2) распарсьте число,
	// TODO: 3) при ошибке парсинга верните ошибку с контекстом через wrapping,
	// TODO: 4) проверьте диапазон реалистичности и верните понятную ошибку, если он нарушен.

	raw := s // TODO: Используйте исходную строку в сообщении об ошибке парсинга (wrapping).
	s = strings.TrimSpace(s)

	age, err := strconv.Atoi(s)
	if err != nil {
		// TODO: Оберните ошибку парсинга, сохранив исходную причину.
		return 0, fmt.Errorf("parse age %q: %w", raw, err)
	}

	// TODO: Добавьте проверки диапазона через ранние возвраты (guard clauses).
	if age < 0 || age > 150 {
		return 0, fmt.Errorf("incorrect age %d", age)
	}
	return age, nil
}
