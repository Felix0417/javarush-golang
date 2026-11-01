package main

import "fmt"

type Pair struct {
	A int
	B int
}

// Sum должен быть методом с получателем-значением (Pair, не *Pair).
func (p Pair) Sum() int {
	// TODO: Реализуйте вычисление результата по полям структуры Pair.
	// TODO: Метод должен возвращать корректное значение для Pair{A: 3, B: 4}.
	return p.A + p.B
}

func main() {
	pair := Pair{A: 3, B: 4}
	pairPtr := &pair

	fmt.Println(pair.Sum(), pairPtr.Sum())
}
