package main

import (
	"bufio"
	"fmt"
	"os"
)

func addIfPositive(dst []int, x int) []int {
	// TODO: Реализуйте функцию: если x <= 0 — вернуть dst без изменений,
	// TODO: если x > 0 — добавить x в dst через append и вернуть результат.
	if x <= 0 {
		return dst
	}
	return append(dst, x)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	positiveReadings := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)

		// TODO: Важно: результат addIfPositive нужно присваивать обратно в positiveReadings,
		// TODO: потому что append может вернуть новый заголовок слайса.
		positiveReadings = addIfPositive(positiveReadings, x)
	}

	// TODO: Выведите все числа из positiveReadings через пробел в одной строке.
	// TODO: Если positiveReadings пустой — выведите только перевод строки.
	fmt.Fprintln(out, positiveReadings)
}
