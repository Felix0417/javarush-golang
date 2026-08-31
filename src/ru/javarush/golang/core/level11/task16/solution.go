package main

import "fmt"

func main() {
	s := make([]int, 0, 1)

	// TODO: Реализуйте сценарий, в котором после роста s происходит смена backing array, и изменения через t больше не влияют на s.
	s = append(s, 1)
	t := s
	// TODO: Добавьте необходимые операции со слайсами (включая append и изменение элемента) в правильном порядке.
	s = append(s, 2)
	t[0] = 99

	fmt.Println(s)
	fmt.Println(t)
}
