package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numberToken string
	fmt.Fscan(os.Stdin, &numberToken)

	n, err := strconv.Atoi(numberToken)

	// TODO: Реализуйте ветвление строго по err (err == nil / err != nil) и корректный вывод по формату задачи.
	if err == nil {
		fmt.Println("OK")
		// TODO: Выведите распознанное целое число n во второй строке.
		fmt.Println(n)
	} else {
		fmt.Println("ERROR")
		// TODO: Выведите само значение err отдельной строкой (через fmt.Println(err) или %v), без анализа err.Error().
		// TODO: Выведите тип err с помощью fmt.Printf("%T\n", err).
		fmt.Println(err)
		fmt.Printf("%T\n", err)
	}
}
