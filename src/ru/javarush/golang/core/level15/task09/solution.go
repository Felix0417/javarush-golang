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

	sums := make(map[string]int)

	for i := 0; i < n; i++ {
		var category string
		var amount int
		fmt.Fscan(in, &category, &amount)

		// TODO: Реализуйте корректное накопление суммы по категориям (повторы должны увеличивать сумму, а не затирать предыдущее значение).
		if v, ok := sums[category]; ok {
			sums[category] = v + amount
			continue
		}
		sums[category] = amount
	}

	total := 0
	for _, v := range sums {
		total += v
	}

	fmt.Fprintln(out, total)
}
