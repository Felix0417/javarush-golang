package main

import (
	"bufio"
	"fmt"
	"os"
)

func setAt(nums []int, index int, value int) bool {
	if index < 0 || index >= len(nums) {
		return false
	}

	// TODO: Реализуйте изменение элемента "на месте" через указатель:
	// TODO: 1) возьмите адрес элемента nums[index] через &nums[index]
	// TODO: 2) запишите value через разыменование указателя (*p = value)
	// TODO: Важно: при корректном индексе функция должна вернуть true.
	addr := &nums[index]
	*addr = value

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	trackIDs := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &trackIDs[i])
	}

	var indexToChange, newValue int
	fmt.Fscan(in, &indexToChange, &newValue)

	if !setAt(trackIDs, indexToChange, newValue) {
		fmt.Fprintln(out, "ERROR")
		return
	}

	for i, v := range trackIDs {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}
