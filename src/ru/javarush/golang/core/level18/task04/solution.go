package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	printClosure(n)
	printArgs(n)
}

func printClosure(n int) {
	fmt.Print("closure: ")
	defer fmt.Println()

	for i := 0; i < n; i++ {
		// TODO: Зарегистрируйте defer для печати цифры так, чтобы проявился эффект замыкания (как требуется в задаче).
		defer func() {
			fmt.Print(i)
		}()
	}
}

func printArgs(n int) {
	fmt.Print("args: ")
	defer fmt.Println()

	for i := 0; i < n; i++ {
		// TODO: Зарегистрируйте defer для печати цифры через аргумент анонимной функции (как требуется в задаче).
		defer func(v int) {
			fmt.Print(v)
		}(i)
	}
}
