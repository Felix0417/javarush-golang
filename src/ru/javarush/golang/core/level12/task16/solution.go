package main

import (
	"bufio"
	"fmt"
	"os"
)

func SnapshotLastK(nums []int, k int) []int {
	if k <= 0 {
		return make([]int, 0)
	}

	if k >= len(nums) {
		// TODO: Верните независимую копию всего nums (новый backing array), а не исходный слайс.
		copySlice := make([]int, len(nums))
		copy(copySlice, nums)
		return copySlice
	}

	start := len(nums) - k
	// TODO: Верните независимую копию последних k элементов через make+copy, а не под-слайс, который делит память с nums.
	copySlice := make([]int, k)
	copy(copySlice, nums[start:])

	return copySlice
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	var k int
	fmt.Fscan(in, &k)

	result := SnapshotLastK(nums, k)

	// Проверяем независимость: меняем result, nums не должен измениться.
	if len(result) > 0 {
		result[0] = 777
	}

	fmt.Println(nums)
	fmt.Println(result)
}
