package main

import "fmt"

type Wallet struct {
	Amount int
}

func (w *Wallet) Add(delta int) {
	// TODO: Реализуйте пополнение кошелька: обновите поле Amount у текущего кошелька на величину delta.
	w.Amount += delta
}

func main() {
	var startAmount int
	var delta int
	fmt.Scan(&startAmount, &delta)

	var wallet Wallet
	wallet.Amount = startAmount

	wallet.Add(delta)

	fmt.Println(wallet.Amount)
}
