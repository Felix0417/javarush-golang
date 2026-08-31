package main

import (
	"bufio"
	"fmt"
	"os"
)

func addScore(scores map[string]int, playerName string, delta int) map[string]int {
	// TODO: Реализуйте обновление карты очков:
	// TODO: 1) Если scores == nil — инициализируйте карту строго через make(map[string]int).
	// TODO: 2) Прибавьте delta к текущему значению scores[playerName] строго через выражение:
	// TODO:    scores[playerName] = scores[playerName] + delta
	// TODO: 3) Обязательно верните (возможно, новую) карту scores.
	if scores == nil {
		scores = make(map[string]int)
	}
	scores[playerName] = scores[playerName] + delta

	return scores
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var scores map[string]int // важно: nil-map, без make

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var playerName string
		var delta int
		fmt.Fscan(in, &playerName, &delta)

		// Важно присваивать результат обратно: карта могла быть nil.
		scores = addScore(scores, playerName, delta)
	}

	var queryName string
	fmt.Fscan(in, &queryName)

	fmt.Fprintln(out, scores[queryName])
	fmt.Fprintln(out, len(scores))

}
