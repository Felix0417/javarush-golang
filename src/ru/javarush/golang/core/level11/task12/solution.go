package main

import (
	"bufio"
	"fmt"
	"os"
)

// FirstString безопасно возвращает первый элемент слайса и флаг успеха.
func FirstString(items []string) (string, bool) {
	// TODO: Реализуйте безопасное получение первого элемента слайса.
	// TODO: Нельзя определять наличие первого элемента через сравнение с nil — проверяйте длину.
	if len(items) == 0 {
		return "", false
	} else if len(items) > 0 {
		return items[0], true
	}
	return "", false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	// Для n=0 items должен остаться nil-слайсом.
	var items []string

	// TODO: Инициализируйте items в зависимости от n так, чтобы при n == 0 он оставался nil-слайсом.
	// TODO: При n > 0 создайте слайс нужной длины и заполните его по индексам.
	if n == 0 {
		items = nil
	} else if n >= 0 {
		items = make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &items[i])
		}
	}

	first, ok := FirstString(items)

	// TODO: Исправьте вывод так, чтобы он отражал именно то, является ли items nil-слайсом.
	fmt.Fprintf(out, "slice_nil=%v\n", items == nil)

	if ok {
		fmt.Fprintf(out, "first=%s\n", first)
	} else {
		fmt.Fprintln(out, "first=<none>")
	}
}
