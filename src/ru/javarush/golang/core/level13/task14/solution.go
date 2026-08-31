package main

import (
	"bufio"
	"fmt"
	"os"
)

func insertSlice(s []int, insertPos int, add []int) []int {
	k := len(add)
	if k == 0 {
		return s
	}
	if insertPos < 0 || insertPos > len(s) {
		return s
	}

	// Увеличиваем длину строго через append + make (по требованиям).
	s = append(s, make([]int, k)...)

	// TODO: Реализуйте вставку блока add в позицию insertPos.
	// TODO: Нужно сдвинуть хвост исходного слайса вправо на k позиций с помощью copy,
	// TODO: затем скопировать элементы add в освободившийся диапазон с помощью copy.
	copy(s[insertPos+k:], s[insertPos:])
	copy(s[insertPos:], add)

	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	warehouseStock := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &warehouseStock[i])
	}

	var k int
	fmt.Fscan(in, &k)

	incomingBatch := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &incomingBatch[i])
	}

	var insertPos int
	fmt.Fscan(in, &insertPos)

	res := insertSlice(warehouseStock, insertPos, incomingBatch)

	for i := 0; i < len(res); i++ {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, res[i])
	}
	fmt.Fprintln(out)
}
