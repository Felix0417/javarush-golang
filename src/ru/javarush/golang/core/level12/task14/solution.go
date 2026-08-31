package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(in, &N)

	src := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &src[i])
	}

	dst := make([]int, 0, N)

	// TODO: Перед вызовом copy обеспечьте, чтобы длина dst была равна N (len(dst) == N).

	dst = dst[:N]
	n := copy(dst, src)

	fmt.Println(n)
	fmt.Println(dst)
}
