package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)

	// По условию словарь нужно создавать именно через make (а не new).
	dict := make(map[string]string)

	for i := 0; i < k; i++ {
		var key, value string
		fmt.Scan(&key, &value)

		// TODO: Добавьте пару key/value в словарь через присваивание dict[key] = value.
		dict[key] = value
	}

	var query string
	fmt.Scan(&query)

	// TODO: Реализуйте поиск query в dict через ok-идиому: value, ok := dict[query].
	// TODO: Если ключ найден — выведите value, иначе выведите "not found".
	value, ok := dict[query]
	if !ok {
		fmt.Println("not found")
		return
	}

	fmt.Println(value)
}
