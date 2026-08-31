package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var D, S, F int
	fmt.Fscan(in, &D, &S, &F)

	// src заданной длины S
	src := make([]int, S)
	for i := 0; i < S; i++ {
		fmt.Fscan(in, &src[i])
	}

	// dst именно длины D (не 0 с capacity)
	dst := make([]int, D)

	// TODO: заполните все элементы dst значением F с помощью цикла for
	for i := 0; i < D; i++ {
		dst[i] = F
	}

	// TODO: выполните копирование через n := copy(dst, src) без создания дополнительных слайсов
	n := copy(dst, src)

	// TODO: выведите ровно две строки: в первой n, во второй — содержимое dst после копирования
	// TODO: учтите, что при D = 0 вторая строка может быть пустой
	fmt.Fprintln(out, n)
	fmt.Fprintln(out, dst)
}
