package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Fscan(os.Stdin, &s)

	b := make([]byte, 0, len(s)+2)

	// TODO: Соберите результат строго в формате "<s>" через []byte и append:
	// TODO: 1) добавьте байт '<'
	// TODO: 2) добавьте байты строки s через append с раскрытием s...
	// TODO: 3) добавьте байт '>'
	// TODO: После этого один раз преобразуйте b в string и выведите.
	b = append(b, '<')
	b = append(b, s...)
	b = append(b, '>')

	fmt.Println(string(b))
}
