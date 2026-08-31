package main

import (
	"bufio"
	"fmt"
	"os"
)

func printSlice(w *bufio.Writer, s []int) {
	for i, v := range s {
		if i > 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, v)
	}
	fmt.Fprintln(w)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	src := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &src[i])
	}

	snapshot := make([]int, len(src))

	// TODO: Скопируйте все элементы из src в snapshot так, чтобы snapshot был независимой копией (без общего backing array с src).
	// TODO: Не используйте присваивание вида snapshot := src (или эквивалент), потому что это создаёт aliasing.

	copy(snapshot, src)

	if len(snapshot) > 0 {
		snapshot[0] = -1
	}

	printSlice(out, src)
	printSlice(out, snapshot)
}
