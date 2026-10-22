package main

import "fmt"

func main() {
	tasks := []string{"buy_milk", "learn_go", "sleep"}

	ptrs := make([]*string, 0, len(tasks))
	// TODO: Соберите ptrs так, чтобы изменение через ptrs демонстрировало ловушку указателя на range-переменную и НЕ меняло tasks.
	for _, v := range tasks {
		ptrs = append(ptrs, &v)
	}

	*ptrs[1] = "HACKED"

	fmt.Println("через указатель:", *ptrs[1])
	fmt.Println("в слайсе:", tasks[1])
}
