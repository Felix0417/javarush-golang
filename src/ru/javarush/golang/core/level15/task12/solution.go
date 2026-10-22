package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	counts := make(map[string]int)
	for i := 0; i < n; i++ {
		var w string
		fmt.Fscan(in, &w)

		// TODO: реализуйте корректный подсчёт частот слов в map (при повторе слова счётчик должен увеличиваться)
		if v, ok := counts[w]; ok {
			counts[w] = v + 1
			continue
		}
		counts[w] = 1
	}

	var queryWord string
	fmt.Fscan(in, &queryWord)

	// TODO: увеличьте каждый счётчик на 10 и обязательно запишите обновлённое значение обратно в map по ключу (не пытайтесь менять только переменную v)
	for k, v := range counts {
		counts[k] = v + 10
	}

	fmt.Printf("%d %d\n", len(counts), counts[queryWord])
}
