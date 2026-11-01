package main

import "fmt"

type ShoppingList struct {
	Items []string
}

func (s *ShoppingList) Add(item string) {
	// TODO: Реализуйте добавление товара в список так, чтобы изменение сохранялось "снаружи":
	// TODO: используйте append и обязательно присвойте результат обратно в s.Items.
	s.Items = append(s.Items, item)
}

func main() {
	var n int
	fmt.Scan(&n)

	list := ShoppingList{
		Items: make([]string, 0, n),
	}

	for i := 0; i < n; i++ {
		var productName string
		fmt.Scan(&productName)
		list.Add(productName)
	}

	for _, item := range list.Items {
		fmt.Println(item)
	}
}
