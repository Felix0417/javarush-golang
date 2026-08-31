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

	s := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	var a, b int
	fmt.Fscan(in, &a, &b)

	// TODO: Создайте view как под-слайс s[a:b] без копирования (общий underlying array).
	// TODO: Создайте snapshot как независимую копию элементов view через make+copy.
	view := s[a:b]
	snapshot := make([]int, len(view))
	copy(snapshot, view)

	// TODO: Если len(view) > 0 — увеличьте view[0] на 1 (не обращайтесь к [0], если view пустой).
	// TODO: Если len(snapshot) > 0 — установите snapshot[0] в -1000 (не обращайтесь к [0], если snapshot пустой).
	if len(view) > 0 {
		view[0] = view[0] + 1
	}
	if len(snapshot) > 0 {
		snapshot[0] = -1000
	}

	fmt.Println(s)
	fmt.Println(view)
	fmt.Println(snapshot)
}
