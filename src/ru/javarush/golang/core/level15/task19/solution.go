package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	// Счётчик частот: ключ — слово, значение — сколько раз встретилось.
	counts := make(map[string]int)

	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(in, &word)

		// TODO: увеличьте частоту слова строго операцией counts[word]++
		counts[word]++
	}

	// Для стабильного вывода нужны ключи и сортировка.
	keys := make([]string, 0, len(counts))

	// TODO: соберите все ключи counts в keys
	for k := range counts {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Fprintln(out, len(counts))

	// TODO: выведите ровно len(counts) строк в лексикографическом порядке слов
	// TODO: формат строки: слово + пробел + количество
	// Сейчас намеренно неверно: порядок нестабилен и количество не выводится.
	for _, k := range keys {
		fmt.Fprintln(out, k, counts[k])
	}
}
