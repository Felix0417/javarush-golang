package main

import (
	"errors"
	"fmt"
	"os"
)

// Sentinel-ошибка: обязана быть единой для случая, когда элементов меньше 3.
var ErrNeedAtLeast3 = errors.New("need at least 3 numbers")

// Third возвращает третий элемент (индекс 2) без паники на коротком вводе.
func Third(nums []int) (int, error) {
	if len(nums) < 3 {
		return 0, ErrNeedAtLeast3
	}

	// TODO: Верните третий элемент (индекс 2) из слайса nums, не допуская panic.
	return nums[2], nil
}

func main() {
	in := os.Stdin

	var n int
	fmt.Fscan(in, &n)

	// Важно: make([]int, n) с отрицательным n вызовет panic, поэтому проверяем заранее.
	if n < 0 {
		fmt.Println("error:", errors.New("n must be non-negative"))
		return
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	v, err := Third(nums)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(v)
}
