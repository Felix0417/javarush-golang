package main

import "fmt"

// Order — модель заказа: ID и признак оплаты.
type Order struct {
	ID   int
	Paid bool
}

// IsPaid — value receiver по требованию: метод ничего не меняет, только читает состояние заказа.
func (o Order) IsPaid() bool {
	// TODO: Реализуйте метод так, чтобы он корректно сообщал, оплачен ли заказ.
	return o.Paid
}

func main() {
	var orderID int
	var paidFlag int
	fmt.Scan(&orderID, &paidFlag)

	order := Order{
		ID: orderID,
		// TODO: Заполните поле Paid, преобразовав paidFlag (0/1) в bool.
		Paid: paidFlag == 1,
	}

	if order.IsPaid() {
		fmt.Println("PAID")
	} else {
		fmt.Println("NOT PAID")
	}
}
