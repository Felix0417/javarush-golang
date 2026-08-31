package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for idx := range a {
		fmt.Fscan(in, &a[idx])
	}

	var i int
	fmt.Fscan(in, &i)

	// TODO: Проверьте корректность индекса i (диапазон [0; len(a))) и удалите ровно один элемент a[i],
	// TODO: используя slices.Delete(a, i, i+1), сохранив результат обратно в a. Если индекс некорректный — не удаляйте.
	if i >= 0 && i < len(a) {
		a = slices.Delete(a, i, i+1)
	}

	out := bufio.NewWriter(os.Stdout)
	for idx, v := range a {
		if idx > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	out.Flush()
}
