package main

import (
	"bufio"
	"fmt"
	"os"
)

// filterRangeNew строит новый слайс, не трогая исходный s (никаких in-place приёмов).
func filterRangeNew(s []int, L, R int) []int {
	out := make([]int, 0, len(s)) // предвыделяем ёмкость, добавляем только через append

	// TODO: Реализуйте фильтрацию: добавляйте в out только элементы s, которые попадают в диапазон [L; R] включительно.
	// TODO: Сохраняйте порядок элементов как в исходном s.
	// TODO: Исходный слайс s менять нельзя (не используйте in-place подходы вроде s[:0]); добавление делайте только через append.

	for _, x := range s {
		// TODO: Добавьте проверку диапазона и append подходящих значений.
		if x >= L && x <= R {
			out = append(out, x)
		}
	}

	return out
}

// Печать: числа через одиночные пробелы, для пустого слайса — пустая строка.
func printSlice(w *bufio.Writer, s []int) {
	for i, x := range s {
		if i > 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, x)
	}
	fmt.Fprintln(w)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	var L, R int
	fmt.Fscan(in, &L, &R)

	filtered := filterRangeNew(s, L, R)

	printSlice(out, s)
	printSlice(out, filtered)
}
