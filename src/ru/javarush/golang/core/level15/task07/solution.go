package main

import "fmt"

func main() {
	// Предзаданный справочник скидок (важно: 0 — валидная скидка, не "нет данных")
	discounts := map[string]int{
		"alice": 10,
		"bob":   0,
		"carol": 5,
		"dave":  15,
	}

	var n int
	fmt.Scan(&n)

	missing := 0
	for i := 0; i < n; i++ {
		var login string
		fmt.Scan(&login)

		// TODO: Реализуйте корректную проверку наличия логина в map через ok-идиому (v, ok := discounts[login]).
		// TODO: Учтите, что скидка 0 — это валидное значение (если ok == true, нужно печатать discount: 0).
		// TODO: Увеличивайте missing только если логин отсутствует (ok == false).
		v, ok := discounts[login]

		if !ok {
			fmt.Printf("login: %s no discount data\n", login)
			missing++
			continue
		}

		fmt.Printf("login: %s discount: %d\n", login, v)
	}

	fmt.Printf("missing: %d\n", missing)
}
