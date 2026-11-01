package main

import (
	"bufio"
	"fmt"
	"os"
)

type Item struct {
	ID    int
	Title string
	Done  bool
}

// Items — коллекция задач как слайс-тип (тема лекции про методы на слайс-типах).
type Items []Item

// FindByID делает линейный поиск по ID и возвращает (Item{}, false), если не найдено.
func (it Items) FindByID(id int) (Item, bool) {
	// TODO: Реализуйте линейный поиск по коллекции it и найдите элемент с нужным ID.
	// TODO: Если элемент найден — верните (item, true).
	// TODO: Если элемент не найден — верните строго (Item{}, false).
	for i := range it {
		if it[i].ID == id {
			return it[i], true
		}
	}
	return Item{}, false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var id int
	var title string
	var doneFlag int

	items := make(Items, 0, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &id, &title, &doneFlag)

		items = append(items, Item{
			ID:    id,
			Title: title,
			// TODO: Преобразуйте doneFlag (0/1) в bool и сохраните в поле Done.
			Done: doneFlag == 1,
		})
	}

	var idQuery int
	fmt.Fscan(in, &idQuery)

	found, ok := items.FindByID(idQuery)
	if !ok {
		fmt.Fprint(out, "NOT_FOUND")
		return
	}

	// TODO: Преобразуйте found.Done (bool) обратно в 0/1 для вывода (строго 0 или 1).

	var doneOut int
	if found.Done {
		doneOut = 1
	}

	fmt.Fprintf(out, "FOUND %s %d", found.Title, doneOut)
}
