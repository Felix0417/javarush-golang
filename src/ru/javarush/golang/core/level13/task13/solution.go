package main

import (
	"bufio"
	"fmt"
	"os"
)

func insertAt(s []int, insertIndex int, newCode int) []int {
	if insertIndex < 0 || insertIndex > len(s) {
		return s
	}

	// TODO: Реализуйте вставку newCode в позицию insertIndex.
	// TODO: Увеличение длины сделайте через append (с обязательным присваиванием результата),
	// TODO: а сдвиг элементов вправо — через copy по перекрывающимся диапазонам.
	// TODO: Должны корректно работать случаи insertIndex == 0 и insertIndex == len(s).

	s = append(s, 0)
	copy(s[insertIndex+1:], s[insertIndex:])
	s[insertIndex] = newCode

	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	productCodes := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &productCodes[i])
	}

	var insertIndex, newCode int
	fmt.Fscan(in, &insertIndex, &newCode)

	productCodes = insertAt(productCodes, insertIndex, newCode)

	for i, v := range productCodes {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
}
