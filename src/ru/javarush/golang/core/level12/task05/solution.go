package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	// Под-слайс — это "окно" в исходные данные (общий underlying array).
	view := numbers[1:4]
	// Меняем через view — изменится и исходный numbers.
	view[0] = 999

	fmt.Println("view:", view)
	fmt.Println("numbers:", numbers)
}
