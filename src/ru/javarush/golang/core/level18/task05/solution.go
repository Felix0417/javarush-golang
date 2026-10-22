package main

import "fmt"

func main() {
	demo()
}

func demo() {
	for i := 0; i < 3; i++ {
		fmt.Println("work", i)

		// TODO: Отложите вывод "cleanup i" с помощью defer так, чтобы строки cleanup
		// TODO: печатались после выхода из demo() и строго в обратном порядке: 2, 1, 0.
		defer fmt.Println("cleanup", i)
	}
}
