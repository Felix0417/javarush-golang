package main

import "fmt"

// Item — доменная модель предмета инвентаря.
type Item struct {
	ID    int
	Title string
	Done  bool
}

// Items — коллекция предметов (именованный слайс-тип).
type Items []Item

// Len возвращает количество элементов в коллекции.
func (it Items) Len() int {
	// TODO: Реализуйте корректный подсчёт количества элементов в коллекции Items.
	return len(it)
}

func main() {
	var n int
	fmt.Scan(&n)

	items := make(Items, 0, n)

	for id := 1; id <= n; id++ {
		var title string
		fmt.Scan(&title)

		// TODO: Добавьте в коллекцию items новый Item:
		// TODO: - ID должен быть равен id (от 1 до n)
		// TODO: - Title должен быть равен введённому слову title
		// TODO: - Done должен быть false для всех новых элементов
		items = append(items, Item{ID: id, Title: title})
	}

	fmt.Printf("len: %d\n", items.Len())
}
