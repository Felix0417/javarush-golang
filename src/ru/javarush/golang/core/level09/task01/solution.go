package main

import "fmt"

// greet собирает строку приветствия и ничего не печатает.
func greet(name string) string {
	// TODO: Реализуйте формирование строки приветствия в требуемом формате и верните её. Здесь не должно быть печати в stdout.
	return "Hello, " + name + "!"
}

// printGreeting печатает одну строку, используя результат greet.
func printGreeting(name string) {
	// TODO: Убедитесь, что печатается ровно одна строка приветствия, используя результат greet(name), без лишних пробелов/пустых строк.
	fmt.Println(greet(name))
}

func main() {
	var name string
	fmt.Scan(&name)
	printGreeting(name)
}
