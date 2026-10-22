package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TaskID int

// ParseTaskID парсит "номер заявки" и проверяет инвариант: id строго положительный.
func ParseTaskID(s string) (TaskID, error) {
	// TODO: Нормализуйте ввод через strings.TrimSpace.
	s = strings.TrimSpace(s)

	// TODO: Распарсите число через strconv.Atoi и корректно обработайте ошибку парсинга.
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	// TODO: Проверьте инвариант: id должен быть строго положительным (id > 0).
	// TODO: Если инвариант нарушен — верните ошибку.
	// TODO: Если всё корректно — верните TaskID(n), nil.
	if n <= 0 {
		return 0, errors.New("not implemented")
	}
	return TaskID(n), nil
}

func main() {
	var rawID string
	fmt.Fscan(os.Stdin, &rawID)

	id, err := ParseTaskID(rawID)
	if err != nil {
		fmt.Print("ERROR")
		return
	}

	fmt.Printf("OK %d", id)
}
