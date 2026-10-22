package main

import (
	"fmt"
	"os"
	"strconv"
)

// parseOptionalInt разбирает qtyToken:
// "-" -> (nil, nil), число -> (*int, nil), иначе -> ошибка.
func parseOptionalInt(qtyToken string) (*int, error) {
	if qtyToken == "-" {
		return nil, nil
	}

	n, err := strconv.Atoi(qtyToken)
	if err != nil {
		return nil, err
	}

	// TODO: Верните указатель на разобранное число (а не nil), чтобы количество сохранялось в списке.
	return &n, nil
}

// addItem добавляет элемент строго через append с сохранением результата.
func addItem(names []string, qtys []*int, name string, qty *int) ([]string, []*int) {
	names = append(names, name)

	// TODO: Добавьте количество в параллельный слайс qtys; qty может быть nil и это валидное значение.
	qtys = append(qtys, qty)

	return names, qtys
}

func main() {
	in := os.Stdin

	var n int
	fmt.Fscan(in, &n)

	names := make([]string, 0, n)
	qtys := make([]*int, 0, n)

	for i := 0; i < n; i++ {
		var name, qtyToken string
		fmt.Fscan(in, &name, &qtyToken)

		qty, err := parseOptionalInt(qtyToken)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		names, qtys = addItem(names, qtys, name, qty)
	}

	for i := 0; i < n; i++ {
		if qtys[i] == nil {
			fmt.Printf("%d) %s (без количества)\n", i+1, names[i])
			continue
		}
		fmt.Printf("%d) %s %d\n", i+1, names[i], *qtys[i])
	}
}
