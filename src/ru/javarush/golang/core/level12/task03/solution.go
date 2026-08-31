package main

import "fmt"

func main() {
	// По условию: исходный слайс должен быть создан только так: len=0, cap=5.
	s := make([]int, 0, 5)

	var b int
	fmt.Scan(&b)

	// TODO: Реализуйте проверку, что b находится в диапазоне [0, cap(s)] до выполнения нарезки, чтобы избежать panic.
	// TODO: Если b корректно — сформируйте слайс t нарезкой (без индексирования элементов) и выведите:
	// TODO: 1) сам t через fmt.Println, 2) len=<len(t)>, 3) cap=<cap(t)> (ровно 3 строки).
	// TODO: Если b некорректно — выведите ровно 3 строки: [] затем len=0 затем cap=0.

	// Базовый вывод-заглушка (сейчас программа всегда ведёт себя как при некорректном b).
	if b < 0 || b > cap(s) {
		fmt.Println([]int{})
		fmt.Println("len=0")
		fmt.Println("cap=0")
		return
	}

	t := s[:b]
	fmt.Println(t)
	fmt.Printf("len=%d\n", len(t))
	fmt.Printf("cap=%d\n", cap(t))
}
