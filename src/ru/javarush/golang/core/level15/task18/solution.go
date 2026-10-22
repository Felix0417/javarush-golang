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

	// Словарь храним как set: ключ есть -> фраза разрешена.
	dictionary := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		var phrase string
		fmt.Fscan(in, &phrase)

		// TODO: Добавьте phrase в dictionary как в множество (set) через map[string]struct{}.
		dictionary[phrase] = struct{}{}
	}

	var q int
	fmt.Fscan(in, &q)

	for i := 0; i < q; i++ {
		var query string
		fmt.Fscan(in, &query)

		// TODO: Реализуйте проверку наличия query в dictionary строго через ok-идиому.
		// TODO: Выведите "YES", если query есть в словаре, иначе "NO". Обрабатывайте запросы потоково.
		_, ok := dictionary[query]
		if ok {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
