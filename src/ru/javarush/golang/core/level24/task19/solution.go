package main

/*
Вы делаете учебный “кошелёк”: у него есть balance, его можно посмотреть, а ещё его можно “обнулить” (например, для
тестовой сессии). И тут появляется важное правило Go: переменная‑значение может вызвать метод с pointer receiver
напрямую (если она адресуемая), но интерфейс гораздо строже — в интерфейс “влезают” только те методы, которые есть у
конкретного типа.

Создайте тип Wallet с полем balance int. Опишите два интерфейса: Resetter (Reset()) и Balancer (Balance() int).
Реализуйте Balance() на value receiver (Wallet), а Reset() — на pointer receiver (*Wallet), чтобы Reset обнулял balance.

Напишите две функции: UseReset(r Resetter) и UseBalance(b Balancer). Первая вызывает Reset, вторая печатает текущий баланс.

В main создайте переменную w := Wallet{balance: 100}. Программа должна:
вывести баланс через UseBalance(w),
вывести баланс через UseBalance(&w),
вызвать метод w.Reset() напрямую (это должно работать из-за адресуемости переменной),
вызвать UseReset(&w),
вывести итоговый баланс (должен стать 0).
*/

import "fmt"

type Wallet struct {
	balance int
}

type Resetter interface {
	Reset()
}

type Balancer interface {
	Balance() int
}

func (w Wallet) Balance() int {
	// TODO: Верните текущий баланс кошелька (значение поля balance), не изменяя состояние.
	return w.balance
}

func (w *Wallet) Reset() {
	// TODO: Обнулите баланс кошелька (измените поле balance).
	w.balance = 0
}

func UseReset(r Resetter) {
	// TODO: Вызовите метод Reset() у переданного интерфейса.
	r.Reset()
}

func UseBalance(b Balancer) {
	// TODO: Напечатайте текущий баланс через fmt.Println(b.Balance()).
	fmt.Println(b.Balance())
}

func main() {
	w := Wallet{balance: 100}

	// TODO: Реализуйте сценарий из условия:
	// TODO: 1) Выведите баланс через UseBalance(w)
	// TODO: 2) Выведите баланс через UseBalance(&w)
	// TODO: 3) Вызовите w.Reset() напрямую
	// TODO: 4) Вызовите UseReset(&w)
	// TODO: 5) Выведите итоговый баланс (после сброса)
	UseBalance(w)
	UseBalance(&w)
	w.Reset()
	UseReset(&w)
	fmt.Println(w.Balance())

}
