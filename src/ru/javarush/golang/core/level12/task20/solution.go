package main

import (
	"bufio"
	"fmt"
	"os"
)

func firstNCopy(src []int, n int) []int {
	if n < 0 {
		n = 0
	}
	if n > len(src) {
		n = len(src)
	}

	// TODO: Реализуйте создание независимой копии первых n элементов src.
	// TODO: Результат должен иметь длину n и не должен разделять backing array с src.
	// TODO: Используйте выделение памяти под новый слайс и копирование данных.
	dst := make([]int, n)
	copy(dst, src[:n])
	return dst
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var m, n int
	fmt.Fscan(in, &m, &n)

	orig := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &orig[i])
	}

	copySlice := firstNCopy(orig, n)

	// Мутации только копии: orig не должен меняться ни при записи, ни при append.
	if len(copySlice) > 0 {
		copySlice[0] = 999
		copySlice = append(copySlice, 777)
	}

	fmt.Printf("orig: %v\n", orig)
	fmt.Printf("copy: %v\n", copySlice)
}
