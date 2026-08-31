package main

import "fmt"

// repeat выполняет action ровно n раз; цикл должен быть именно здесь, а не в main.
func repeat(n int, action func()) {
	for i := 0; i < n; i++ {
		// TODO: Вызовите action() внутри цикла, чтобы действие выполнилось ровно n раз.
		action()
	}
}

func main() {
	var separatorLength int
	fmt.Scan(&separatorLength)

	// Печатаем по одному символу '=' через анонимную функцию.
	repeat(separatorLength, func() {
		fmt.Print("=")
	})

	// Ровно один перевод строки в конце, даже если separatorLength <= 0.
	fmt.Print("\n")
}
