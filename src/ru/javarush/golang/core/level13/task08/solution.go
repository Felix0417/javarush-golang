package main

import (
	"bufio"
	"fmt"
	"os"
)

func deleteRange(words []string, i, j int) []string {
	// Нормализуем границы в безопасный диапазон.
	if i < 0 {
		i = 0
	}
	if j > len(words) {
		j = len(words)
	}

	// Пустой диапазон — ничего не меняем.
	if i >= j {
		return words
	}

	oldLen := len(words)

	// TODO: Реализуйте удаление полуинтервала [i:j) "на месте":
	// TODO: закройте "дырку" через copy, вычислите новую длину, очистите освобождённый хвост через clear.
	// TODO: верните укороченный слайс без использования append и без пакета slices.
	moved := copy(words[i:], words[j:])
	newLen := i + moved
	clear(words[newLen:oldLen])

	return words[:newLen]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	words := make([]string, n)
	for k := 0; k < n; k++ {
		fmt.Fscan(in, &words[k])
	}

	var i, j int
	fmt.Fscan(in, &i, &j)

	words = deleteRange(words, i, j)

	for k, w := range words {
		if k > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, w)
	}
	fmt.Fprintln(out)
}
