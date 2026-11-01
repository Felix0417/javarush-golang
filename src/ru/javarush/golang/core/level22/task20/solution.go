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

// Именованный тип-слайс, чтобы дать ему методы.
type Items []Item

func (it *Items) Add(x Item) {
	// TODO: Реализуйте добавление элемента в конец слайса через append так,
	// чтобы изменение длины сохранялось снаружи (нужна запись результата в *it).
	*it = append(*it, x)
}

func (it *Items) DeleteByID(id int) bool {
	// TODO: Реализуйте удаление по id:
	// 	- удалить только первое совпадение
	// 	- сохранить порядок остальных элементов (stable remove)
	// 	- вернуть true если удалили, иначе false
	//
	// 	Текущая реализация удаляет элемент, но НЕ сохраняет порядок.
	s := *it
	for i := 0; i < len(s); i++ {
		if s[i].ID != id {
			continue
		}

		copy(s[i:], s[i+1:])
		*it = s[:len(s)-1]
		s[len(s)-1] = Item{}
		return true
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m int
	fmt.Fscan(in, &m)

	var items Items

	for i := 0; i < m; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		if cmd == "ADD" {
			var id int
			var title string
			fmt.Fscan(in, &id, &title)

			items.Add(Item{
				ID:    id,
				Title: title,
				Done:  false, // по условию
			})
		} else { // "DEL"
			var id int
			fmt.Fscan(in, &id)
			items.DeleteByID(id)
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprintf(out, "len: %d\n", len(items))
	for _, x := range items {
		fmt.Fprintf(out, "%d %s\n", x.ID, x.Title)
	}
}
