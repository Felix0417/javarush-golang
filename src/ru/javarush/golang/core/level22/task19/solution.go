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

type Items []Item

func (it Items) MarkDone(id int) bool {
	// TODO: Реализуйте поиск задачи по ID и пометьте задачу выполненной так,
	// чтобы изменение было видно в main (нельзя менять только копию элемента).
	for i := range it {
		if it[i].ID == id {
			it[i].Done = true
			return true
		}
	}
	return false
}

func doneFlag(done bool) int {
	if done {
		return 1
	}
	return 0
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var id int
	var title string
	var flag int

	items := make(Items, n)
	for i := 0; i < n; i++ {

		fmt.Fscan(in, &id, &title, &flag)

		items[i] = Item{
			ID:    id,
			Title: title,
			Done:  flag == 1,
		}
	}

	var idToMark int
	fmt.Fscan(in, &idToMark)

	ok := items.MarkDone(idToMark)

	fmt.Printf("marked: %v\n", ok)
	for _, it := range items {
		fmt.Printf("%d %s %d\n", it.ID, it.Title, doneFlag(it.Done))
	}
}
