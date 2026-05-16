package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// По условию: при n == 0 ничего больше не читаем.
	if n == 0 {
		fmt.Println("0 0")
		return
	}

	var best int    // лучший результат (максимум), нужен после цикла
	var updates int // сколько раз ставился новый рекорд, нужен после цикла

	for i := 0; i < n; i++ {
		var x int // текущая попытка — область видимости минимальная
		fmt.Scan(&x)

		// Первая попытка всегда ставит рекорд.
		if i == 0 {
			best = x
			updates = updates + 1
			continue
		}

		if x > best {
			best = x
			updates = updates + 1
		}
	}

	fmt.Println(best, updates)
}
