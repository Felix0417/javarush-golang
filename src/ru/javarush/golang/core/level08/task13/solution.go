package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mask := uint(1) << uint(n)
	// TODO: Вычислите маску включения ровно одного бита по номеру n (0..15) с помощью битового сдвига (тип результата должен быть uint).

	fmt.Printf("%d %016b", mask, mask)
}
