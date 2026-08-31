package main

import "fmt"

// isPrime возвращает true, если n — простое число.
func isPrime(n int) bool {

	// По условию все n <= 1 (включая отрицательные) — не простые.
	if n <= 1 {
		return false
	}

	// TODO: Реализуйте проверку простоты числа по условию задачи (поиск делителей циклом for до i*i <= n).

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var candidateNumber int
	fmt.Scan(&candidateNumber)

	if isPrime(candidateNumber) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
