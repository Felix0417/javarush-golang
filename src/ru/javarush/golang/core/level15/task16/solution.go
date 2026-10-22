package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	scores := make(map[string]int)

	for i := 0; i < n; i++ {
		var name string
		var pts int
		fmt.Fscan(in, &name, &pts)

		// TODO: По условию очки нужно суммировать для каждого игрока (а не перезаписывать).
		scores[name] += pts
	}

	// Нельзя полагаться на порядок range по map — собираем ключи в отдельный слайс.
	names := make([]string, 0, len(scores))
	for name := range scores {
		names = append(names, name)
	}

	slices.SortFunc(names, func(a, b string) int {
		// TODO: Реализуйте компаратор для рейтинга:
		// TODO: 1) по убыванию итоговых очков (scores[name])
		// TODO: 2) при равенстве очков — по имени в алфавитном порядке (возрастание)
		// TODO: Важно: строго соблюдайте контракт slices.SortFunc (отриц./положит./0).
		if scores[a] != scores[b] {
			if scores[a] < scores[b] {
				return 1
			}
			return -1
		}

		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	for _, name := range names {
		fmt.Fprintf(out, "%s %d\n", name, scores[name])
	}
}
