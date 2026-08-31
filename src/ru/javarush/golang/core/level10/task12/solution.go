package main

import (
	"fmt"
	_ "math" // blank import по условию (без прямого использования)
	"strconv"
	str "strings" // алиас нужен, потому что в main есть переменная strings
)

func main() {
	var strings string
	fmt.Scanf("%[^\n]", &strings)

	// TODO: Получите значение task из введённой строки. Используйте TrimSpace через алиас str.
	task := str.TrimSpace(strings)

	// TODO: Вычислите id по формуле из условия (в строковом виде). Используйте strconv.Itoa.
	id := strconv.Itoa(len(task))

	// TODO: Проверьте, что вывод ровно из двух строк и строго в нужном формате.
	fmt.Println("started.")
	fmt.Printf("id:%s task:%s\n", id, task)
}
