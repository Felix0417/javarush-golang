package main

import (
	"bufio"
	"fmt"
	"os"
)

func FindIndex(xs []int, target int) (int, bool) {
	// TODO: Реализуйте поиск первого (минимального) индекса, по которому встречается target в xs.
	// TODO: При успехе верните (index, true). Если не найдено — верните (0, false) (или другой индекс, но bool=false обязателен).
	// TODO: Нельзя использовать "магические" значения индекса вроде -1 для сигнала "не найдено".

	if len(xs) == 0 {
		return 0, false
	}

	for i, v := range xs {
		if v == target {
			return i, true
		}
	}

	return 0, false
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	xs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &xs[i])
	}

	var target int
	fmt.Fscan(in, &target)

	idx, ok := FindIndex(xs, target)
	if ok {
		fmt.Printf("FOUND %d\n", idx)
		return
	}
	fmt.Print("NOT_FOUND\n")
}
