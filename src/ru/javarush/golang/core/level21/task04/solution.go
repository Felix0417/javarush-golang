package main

import "fmt"

type Wallet struct {
	Owner   string
	balance int
}

func Deposit(w *Wallet, amount int) {
	// TODO: Реализуйте пополнение кошелька: увеличьте внутренний баланс w.balance на amount.
	w.balance += amount
}

func Withdraw(w *Wallet, amount int) bool {
	// TODO: Реализуйте списание: если средств достаточно — уменьшите w.balance на amount и верните true,
	// TODO: иначе не меняйте баланс и верните false.
	if w.balance < amount {
		return false
	}
	w.balance -= amount
	return true
}

func main() {
	var ownerName string
	var depositAmount, withdrawAmount int
	fmt.Scan(&ownerName, &depositAmount, &withdrawAmount)

	var w Wallet
	w.Owner = ownerName

	Deposit(&w, depositAmount)
	withdrawn := Withdraw(&w, withdrawAmount)

	fmt.Printf("Balance=%d\n", w.balance)
	fmt.Printf("Withdrawn=%t\n", withdrawn)
}
