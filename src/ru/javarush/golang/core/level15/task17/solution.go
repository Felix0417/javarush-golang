package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	uniqueWords := make(map[string]struct{}) // set на map: struct{} не занимает память под значение

	for i := 0; i < n; i++ {
		var word string
		fmt.Scan(&word)

		// TODO: Добавьте считанное слово в множество uniqueWords (set-паттерн на map[string]struct{})
		uniqueWords[word] = struct{}{}
	}

	fmt.Print(len(uniqueWords))
}
