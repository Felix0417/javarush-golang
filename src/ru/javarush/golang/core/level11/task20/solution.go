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

	// TODO: Создайте результирующий слайс evenValues длиной 0 и ёмкостью n (предвыделение памяти).
	evenValues := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)

		// TODO: Добавляйте в evenValues только чётные числа через append, сохраняя исходный порядок.
		if x%2 == 0 {
			evenValues = append(evenValues, x)
		}
	}

	fmt.Fprintf(out, "even_count=%d cap=%d\n", len(evenValues), cap(evenValues))

	for i, v := range evenValues {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprint(out, "\n")
}
