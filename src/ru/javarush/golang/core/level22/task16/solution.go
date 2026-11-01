package main

import (
	"fmt"
)

type Order struct {
	ID   int
	Item string
	Qty  int
}

type item1 struct {
	ID    int
	Title string
	Done  bool
}

type Items []item1

// Человекочитаемый вывод (метод ничего не печатает сам).
func (o Order) String() string {
	// TODO: Реализуйте человекочитаемое представление заказа в формате Order#<id> "<item>" x<qty>.
	// TODO: Поле Item форматируйте через %q. Поля ID/Item/Qty форматируйте явно, без "дампа" всей структуры.
	return fmt.Sprintf("Order#%d %q x%d", o.ID, o.Item, o.Qty)
}

// rawOrder — тот же Order, но без методов: удобно для "сырого" дампа через %+v.
type rawOrder Order

func main() {
	var id, qty int
	var item string

	fmt.Scan(&id, &item, &qty)

	order := Order{ID: id, Item: item, Qty: qty}

	fmt.Println(order)

	// TODO: Выведите "сырой" дамп полей (через преобразование к rawOrder) с форматом %+v, чтобы печатались имена полей.
	fmt.Printf("%+v\n", rawOrder(order))
}
