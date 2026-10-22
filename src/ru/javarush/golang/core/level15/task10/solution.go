package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	clients := make(map[int]string, n)
	for i := 0; i < n; i++ {
		var id int
		var name string
		fmt.Fscan(in, &id, &name)
		clients[id] = name
	}

	// TODO: Найдите минимальный id среди ключей clients обходом map и явным сравнением ключей.
	// TODO: Нельзя опираться на порядок итерации по map: он не гарантирован.
	minID := 0
	minName := ""
	count := 0
	for k, v := range clients {
		if count == 0 {
			minID = k
			minName = v
			count++
		} else if k < minID {
			minID = k
			minName = v
		}
	}

	fmt.Printf("%d %s\n", minID, minName)
}
