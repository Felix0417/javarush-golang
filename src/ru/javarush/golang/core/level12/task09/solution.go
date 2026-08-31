package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}

	// TODO: Сформируйте под-слайс sub через полное выражение среза nums[a:b:c],
	// TODO: чтобы sub содержал элементы с индекса 1 (включительно) по индекс 4 (не включая),
	// TODO: и при этом len(sub) == 3 и cap(sub) == 3.
	sub := nums[1:4:4]

	fmt.Println(sub)
	fmt.Println(len(sub), cap(sub))
}
