package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	// Храним данные по id в map (быстрый доступ),
	// а порядок ввода отдельно — в слайсе (map порядок не гарантирует).
	cards := make(map[int]string, n)
	inputOrder := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var id int
		var title string
		fmt.Fscan(in, &id, &title)

		cards[id] = title

		// TODO: Сохраняйте в inputOrder именно введённый cardID (id), чтобы затем вывести карточки строго в порядке ввода.
		inputOrder = append(inputOrder, id)
	}

	for _, id := range inputOrder {
		// TODO: Убедитесь, что печатается title, соответствующий cardID из inputOrder (а не по индексу цикла).
		fmt.Fprintln(out, id, cards[id])
	}
}
