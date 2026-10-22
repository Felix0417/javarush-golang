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

	baseTrackIDs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &baseTrackIDs[i])
	}

	var a, b, x int
	fmt.Fscan(in, &a, &b, &x)

	// TODO: Сформируйте "безопасный" под-слайс baseTrackIDs[a:b] так, чтобы последующий append
	// TODO: не мог изменить baseTrackIDs (не должно быть общего backing array/capacity).
	frag := make([]int, b-a)

	copy(frag, baseTrackIDs[a:b])

	newList := append(frag, x)

	printSlice(out, baseTrackIDs)
	printSlice(out, newList)
}
