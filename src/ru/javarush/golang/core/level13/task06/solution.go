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

	jobs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &jobs[i])
	}

	// TODO: Удалите последний элемент "на месте":
	// TODO: 1) сначала очистите последний слот backing array через clear на диапазоне длины 1
	// TODO: 2) затем уменьшите длину слайса через reslice (без создания нового слайса)

	newLen := n - 1
	clear(jobs[newLen:])
	jobs = jobs[:newLen]

	out := bufio.NewWriter(os.Stdout)
	for i, s := range jobs {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, s)
	}
	out.Flush()
}
