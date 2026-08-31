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

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	var k, x int
	fmt.Fscan(in, &k, &x)

	// TODO: Сформируйте prefix из первых k элементов так, чтобы cap(prefix) == len(prefix),
	// TODO: и append в prefix не мог изменить исходный слайс nums.
	prefix := nums[:k:k]

	// TODO: Выполните append и обязательно сохраните результат обратно в prefix.
	prefix = append(prefix, x)

	fmt.Println(nums)
	fmt.Println(prefix)
}
