package main

import (
	"errors"
	"fmt"
)

// Sentinel-ошибка для ситуации "товар не найден".
var ErrItemNotFound = errors.New("item not found")

// GetItem ищет товар по id.
// TODO: Если товара нет — верните ошибку, которая wrap'ит ErrItemNotFound через %w и содержит контекст с id.
func GetItem(items map[int]string, id int) (string, error) {
	name, ok := items[id]

	// TODO: Здесь нужно вернуть sentinel-ошибку через wrapping, чтобы её можно было распознать через errors.Is.
	if !ok {
		return "", fmt.Errorf("item %d: %w", id, ErrItemNotFound)
	}

	return name, nil
}

func main() {
	// "Склад" создаём прямо в main: вводим только id для поиска.
	items := map[int]string{
		1: "Keyboard",
		2: "Mouse",
		3: "Monitor",
	}

	var id int
	fmt.Scan(&id)

	name, err := GetItem(items, id)
	if err == nil {
		fmt.Println("ITEM:", name)
		return
	}

	// TODO: Ситуацию "не найдено" нужно распознавать через errors.Is(err, ErrItemNotFound),
	// а не через прямое сравнение ошибок.
	if errors.Is(err, ErrItemNotFound) {
		fmt.Println("NOT_FOUND")
		return
	}

	fmt.Println("ERROR:", err)
}
