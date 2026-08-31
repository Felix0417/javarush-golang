package main

import (
	"bufio"
	"fmt"
	"os"

	"solution/calc"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var op string
	var a, b int
	var debugInt int

	fmt.Fscan(in, &op, &a, &b, &debugInt)

	res, ok := calc.Apply(op, a, b)
	if !ok {
		fmt.Println("unknown operation")
		return
	}

	// По требованию: debug определяется строго так.
	debug := (debugInt == 1)
	fmt.Println(calc.Format(res, debug))
}
