package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	if n == 0 {
		// По условию: для n=0 печатаем только перевод строки и ничего не читаем дальше.
		fmt.Fprintln(out)
		return
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	// TODO: Отсортируйте nums по возрастанию стандартной библиотекой (sort.Ints или slices.Sort), in-place (над этим же слайсом), без использования map.
	sort.Ints(nums)

	// Вывод без лишних пробелов: один пробел между числами + перевод строки в конце.
	fmt.Fprint(out, nums[0])
	for i := 1; i < len(nums); i++ {
		fmt.Fprint(out, " ", nums[i])
	}
	fmt.Fprintln(out)
}
