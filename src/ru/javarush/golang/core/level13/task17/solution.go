package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	// TODO: Сделайте независимую копию слайса nums (не alias).
	// TODO: Используйте slices.Clone (для этого нужно добавить импорт "slices").
	copyNums := slices.Clone(nums)

	// TODO: Убедитесь, что меняется только копия, а оригинальный nums остаётся без изменений.
	if n > 0 {
		copyNums[0] += 1000
	}

	writeSlice(out, nums)
	fmt.Fprintln(out)
	writeSlice(out, copyNums)
	fmt.Fprintln(out)
}

func writeSlice(w *bufio.Writer, s []int) {
	for i, v := range s {
		if i > 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, v)
	}
}
