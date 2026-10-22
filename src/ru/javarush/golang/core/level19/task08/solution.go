package main

import "fmt"

func addDelta(m map[string]int, key string, delta int) {
	// TODO: Увеличьте значение по ключу key на delta в переданной карте.
	m[key] = m[key] + delta
}

func resetBad(m map[string]int) {
	// TODO: Создайте внутри новую карту через make и запишите в неё "new" -> 999, ничего не возвращая.
	// TODO: Убедитесь, что это не изменяет исходную карту снаружи функции.
	m = make(map[string]int)
	m["new"] = 999
}

func clearAll(m map[string]int) {
	// TODO: Удалите все ключи из переданной карты (через range по ключам и delete), не создавая новую карту.
	for k := range m {
		delete(m, k)
	}
}

func main() {
	var keyA, keyB string
	fmt.Scan(&keyA, &keyB)

	counts := map[string]int{
		keyA: 1,
		keyB: 2,
	}

	addDelta(counts, keyA, 10)
	resetBad(counts)

	fmt.Println(counts[keyA])
	fmt.Println(counts[keyB])
	fmt.Println(counts["new"])

	clearAll(counts)
	fmt.Println(len(counts))
}
