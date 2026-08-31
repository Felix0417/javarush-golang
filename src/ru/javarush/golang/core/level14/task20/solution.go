package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	// По условию читаем ровно две строки именно через ReadString('\n').
	expectedRaw, _ := r.ReadString('\n')
	actualRaw, _ := r.ReadString('\n')

	// Norm строго через TrimSpace от raw-строк.
	expectedNorm := strings.TrimSpace(expectedRaw)
	actualNorm := strings.TrimSpace(actualRaw)

	if expectedNorm == actualNorm {
		fmt.Println("match")
		return
	}

	fmt.Printf("expectedRaw=<%q>\n", expectedRaw)
	fmt.Printf("actualRaw=<%q>\n", actualRaw)
	fmt.Printf("expectedNorm=<%q>\n", expectedNorm)
	fmt.Printf("actualNorm=<%q>\n", actualNorm)

	idx, expB, actB := firstDiff(expectedNorm, actualNorm)
	fmt.Printf("firstDiff=%d %d %d\n", idx, expB, actB)
}

func firstDiff(expectedNorm, actualNorm string) (idx int, expByte int, actByte int) {
	// TODO: Реализуйте поиск первого отличия на уровне байтов в нормализованных строках.
	// TODO: Нужно сравнивать именно []byte(expectedNorm) и []byte(actualNorm).
	// TODO: Верните индекс отличия, байт expected (или -1 если expected закончилась), байт actual (или -1 если actual закончилась).
	// TODO: Учтите случай, когда одна строка является префиксом другой (отличие только в длине).
	expectedBytes := []byte(expectedNorm)
	actualBytes := []byte(actualNorm)
	n := len(expectedBytes)
	if len(actualBytes) < n {
		n = len(actualBytes)
	}

	for i := 0; i < n; i++ {
		if expectedBytes[i] != actualBytes[i] {
			return i, int(expectedBytes[i]), int(actualBytes[i])
		}
	}
	i := n
	exp := -1
	act := -1
	if i < len(expectedBytes) {
		exp = int(expectedBytes[i])
	}
	if i < len(actualBytes) {
		act = int(actualBytes[i])
	}
	return i, exp, act
}
