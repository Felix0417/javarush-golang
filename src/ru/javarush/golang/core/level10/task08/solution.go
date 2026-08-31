package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var path string
		fmt.Scan(&path)

		lastSlash := strings.LastIndex(path, "/")

		// TODO: По правилу "последнего сегмента" выведите ожидаемое имя пакета.
		// TODO: Используйте strings.LastIndex для поиска последнего символа "/".
		// TODO: Если "/" нет, нужно вывести весь путь; если есть — вывести только последний сегмент.
		if lastSlash == -1 {
			fmt.Println(path)
		} else {
			fmt.Println(path[lastSlash+1:])
		}
	}
}
