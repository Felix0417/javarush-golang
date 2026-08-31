package main

import "fmt"

func main() {
	var nickname string
	fmt.Scan(&nickname) // один токен, без пробелов

	bytesLen := len(nickname) // длина в байтах (как уйдёт по сети)

	runesLen := 0
	for range nickname { // range по строке идёт по рунам (Unicode-символам)
		// TODO: Увеличьте счётчик, чтобы посчитать количество рун (Unicode-символов) обходом строки через for range.
		runesLen++
	}

	fmt.Printf("%d %d\n", bytesLen, runesLen)
}
