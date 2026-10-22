package main

import (
	"bufio"
	"fmt"
	"os"
)

func buildSquaresPtrs(n int) []*int {
	if n <= 0 {
		// Важно: нужен именно пустой слайс длины 0, а не nil.
		return make([]*int, 0)
	}

	// Предвыделяем место под n указателей.
	res := make([]*int, 0, n)
	for i := 1; i <= n; i++ {
		// TODO: Вычислите квадрат числа i и сохраните результат в переменную sq.
		sq := i * i

		// TODO: Убедитесь, что в результирующем слайсе нет nil-указателей и что значения корректны.
		if &sq == nil {
			continue
		}
		res = append(res, &sq)
	}

	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	ptrs := buildSquaresPtrs(n)
	for i, p := range ptrs {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, *p)
	}
	fmt.Fprintln(out)
}
