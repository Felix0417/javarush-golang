package main

import "fmt"

func main() {
	var budget int
	var purchase int
	var cashback int

	fmt.Scan(&budget, &purchase, &cashback)

	rest := budget

	rest -= purchase
	rest += cashback

	fmt.Println(rest)
}
