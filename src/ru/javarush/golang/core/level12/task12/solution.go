package main

import "fmt"

func Split(s []int, k int) (left []int, right []int) {
	if k < 0 {
		k = 0
	}
	if k > len(s) {
		k = len(s)
	}

	// TODO: Сделайте left "безопасным" view на первые k элементов: ограничьте ёмкость так,
	// чтобы append(left, ...) не мог перезаписать элементы правой части в том же backing array.
	left = s[:k:k]

	right = s[k:]
	return left, right
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	left, right := Split(s, 3)

	left = append(left, 100, 200)

	fmt.Println(s)
	fmt.Println(right)
	fmt.Println(left)
}
