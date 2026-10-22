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

	var wordCount int
	fmt.Fscan(in, &wordCount)

	counts := make(map[string]int)
	keys := make([]string, 0, wordCount)

	for i := 0; i < wordCount; i++ {
		var word string
		fmt.Fscan(in, &word)

		// TODO: перед выводом нужно собрать уникальные слова в отдельный слайс []string
		if _, ok := counts[word]; !ok {
			keys = append(keys, word)
		}
		// TODO: увеличьте счётчик строго через выражение counts[word]++
		counts[word]++
	}

	// TODO: отсортируйте слайс уникальных ключей по алфавиту и выводите в этом порядке
	// TODO: вывод должен быть ровно по одной строке на каждое уникальное слово: "word count"
	slices.Sort(keys)

	for _, word := range keys {
		fmt.Fprintln(out, word, counts[word])
	}
}
