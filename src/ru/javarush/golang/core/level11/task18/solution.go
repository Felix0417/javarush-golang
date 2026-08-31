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

	// TODO: Создайте слайс values строго через make([]int, 0, n), чтобы len был 0, а cap был n.
	values := make([]int, 0, n)

	sum := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)

		// TODO: Добавьте x в values только через append (без индексной записи) и посчитайте сумму.
		values = append(values, x)
		sum += x
	}

	// TODO: Выведите в первой строке все числа из values в исходном порядке через пробел.
	for _, v := range values {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	fmt.Fprintf(out, "len=%d cap=%d\n", len(values), cap(values))
	fmt.Fprintln(out, sum)
}
