package main

import "fmt"

func checkAge(age int) string {
	fmt.Println("start")

	// TODO: Отложите печать строки "finish" через defer так, чтобы она печаталась при любом выходе из функции.
	defer fmt.Println("finish")
	// TODO: Реализуйте проверку возраста:
	// TODO: - если age < 0, верните "invalid" (ранним return)
	// TODO: - иначе верните "ok"
	if age < 0 {
		return "invalid"
	}
	return "ok"
}

func main() {
	var age int
	fmt.Scan(&age)

	res := checkAge(age)
	fmt.Printf("result: %s\n", res)
}
