package main

import "fmt"

// split2Digits возвращает десятки и единицы числа n (0..99).
func split2Digits(n int) (int, int) {
	// TODO: Реализуйте выделение десятков и единиц для числа 0..99 и верните их в порядке: десятки, затем единицы.
	return n / 10, n % 10
}

func main() {
	var n int
	fmt.Scan(&n)

	// Важно: множественное присваивание ровно в таком виде по требованию задачи.
	tens, ones := split2Digits(n)

	// Два числа в одной строке через пробел + перевод строки.
	fmt.Println(tens, ones)
}
