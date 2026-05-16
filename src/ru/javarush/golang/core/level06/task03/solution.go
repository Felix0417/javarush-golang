package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("0 0")
		return
	}

	// "Накопители" живут дольше цикла: объявляем снаружи.
	var sum int
	var max int

	for i := 0; i < n; i++ {
		// Текущее число — внутри тела цикла, чтобы показать область видимости.
		var x int
		fmt.Scan(&x)

		sum = sum + x

		if i == 0 {
			max = x
		} else if x > max {
			max = x
		}
	}

	fmt.Println(sum, max)
}
