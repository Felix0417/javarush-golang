package main

import (
	"bufio"
	"fmt"
	"os"
)

// NonNilInts нормализует nil-слайс в empty-слайс.
// Важно: для не-nil нужно вернуть тот же самый слайс (без копирования).
func NonNilInts(s []int) []int {
	// TODO: Реализуйте нормализацию: если s == nil, вернуть empty-слайс (len=0, но не nil),
	// TODO: а если s != nil — вернуть исходный слайс без копирования и пересоздания.
	if s == nil {
		return []int{}
	}
	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var nums []int
	if n == -1 {
		// TODO: При n == -1 нужно создать nil-слайс nums (len=0, nil=true).
		nums = nil
	} else if n >= 0 {
		// По требованию: создаём через make и заполняем по индексам (без append).
		nums = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &nums[i])
		}
	} else {
		return
	}

	fmt.Printf("before: nums=%v len=%d nil=%t\n", nums, len(nums), nums == nil)

	normalized := NonNilInts(nums)
	fmt.Printf("after: nums=%v len=%d nil=%t\n", normalized, len(normalized), normalized == nil)
}
