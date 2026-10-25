package main

import (
	"bufio"
	"fmt"
	"os"
)

type Article struct {
	ID   int
	Tags []string
}

func equalStringSlices(a, b []string) bool {
	// TODO: Реализуйте сравнение двух слайсов строк по смыслу:
	// TODO: вернуть true только если длины равны и все элементы на одинаковых индексах совпадают (порядок важен).
	// TODO: учтите, что слайсы длины 0 должны считаться равными независимо от того, nil они или пустые.

	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func sameArticle(a, b Article) bool {
	// TODO: Реализуйте сравнение статей: ID должны совпадать и теги должны совпадать по equalStringSlices.
	// TODO: нельзя сравнивать структуры через == из-за поля-слайса.
	return a.ID == b.ID && equalStringSlices(a.Tags, b.Tags)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var id1, m1 int
	fmt.Fscan(in, &id1, &m1)

	var tags1 []string
	for i := 0; i < m1; i++ {
		var t string
		fmt.Fscan(in, &t)
		tags1 = append(tags1, t)
	}

	var id2, m2 int
	fmt.Fscan(in, &id2, &m2)

	var tags2 []string
	for i := 0; i < m2; i++ {
		var t string
		fmt.Fscan(in, &t)
		tags2 = append(tags2, t)
	}

	a1 := Article{ID: id1, Tags: tags1}
	a2 := Article{ID: id2, Tags: tags2}

	if sameArticle(a1, a2) {
		fmt.Print("SAME")
	} else {
		fmt.Print("DIFF")
	}
}
