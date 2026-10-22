package main

import (
	"bufio"
	"fmt"
	"os"
)

func elementPointers(nums []int) []*int {
	ptrs := make([]*int, len(nums))
	for i := range nums {
		// TODO: Верните указатели именно на элементы исходного слайса nums,
		// TODO: а не на копии значений. Указатели должны соответствовать порядку nums.
		ptrs[i] = &nums[i]
	}
	return ptrs
}

func mulAll(ptrs []*int, k int) {
	// TODO: Умножьте КАЖДОЕ значение по указателю на k через разыменование,
	// TODO: изменяя значения "на месте".
	if len(ptrs) == 0 {
		return
	}
	for i := range ptrs {
		*ptrs[i] *= k
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	var k int
	fmt.Fscan(in, &k)

	ptrs := elementPointers(nums)

	// TODO: После получения ptrs нельзя делать append к nums, чтобы не сломать указатели.
	mulAll(ptrs, k)

	for i, v := range nums {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}
