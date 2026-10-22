package main

import "fmt"

func setFirst(nums []int, v int) {
	// TODO: Если nums не пустой, присвойте nums[0] = v.
	if len(nums) > 0 {
		nums[0] = v
	}
}

func addValue(nums []int, v int) []int {
	// TODO: Добавьте v в nums через append и верните новый слайс.
	nums = append(nums, v)
	return nums
}

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	setFirst(nums, 100)

	addValue(nums, 200)

	fmt.Printf("len: %d\n", len(nums))
	fmt.Printf("nums: %v\n", nums)

	nums = addValue(nums, 300)

	fmt.Printf("len: %d\n", len(nums))
	fmt.Printf("nums: %v\n", nums)
}
