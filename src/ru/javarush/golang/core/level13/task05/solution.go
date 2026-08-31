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

	// Слайс фиксированной длины: len(readings) всегда == n.
	readings := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &readings[i])
	}

	var from, to int
	fmt.Fscan(in, &from, &to)

	// TODO: Обнулите только указанный диапазон элементов через clear, применив clear к под-слайсу нужного участка.
	// Нельзя занулять элементы вручную по одному присваиванием.
	clear(readings[from:to])

	out := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		// По условию допустим завершающий пробел.
		fmt.Fprint(out, readings[i], " ")
	}
	out.Flush()
}
