package main

import "fmt"

func main() {
	var visitCount int
	fmt.Scan(&visitCount)

	visitCountPtr := &visitCount // берём адрес переменной

	*visitCountPtr++

	fmt.Println(visitCount)
}
