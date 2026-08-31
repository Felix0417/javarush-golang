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

	src := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &src[i])
	}

	// TODO: Создайте независимую копию dst: выделите отдельный буфер через make и скопируйте элементы из src с помощью copy.
	dst := make([]int, len(src))
	copy(dst, src)

	if n > 0 {
		dst[0] = 999
	}

	fmt.Println(src)
	fmt.Println(dst)
}
