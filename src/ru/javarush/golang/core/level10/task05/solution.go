package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Импортируется по пути math/rand, но в коде обращаемся по имени пакета: rand
	n := rand.Intn(1)

	// TODO: Выведите значение n одной строкой так, чтобы вывод полностью совпал с ожидаемым (включая перевод строки).
	fmt.Println(n)
}
