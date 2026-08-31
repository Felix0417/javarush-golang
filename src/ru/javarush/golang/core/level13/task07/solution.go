package main

import (
	"bufio"
	"fmt"
	"os"
)

func removeAtStable(guestNames []string, removeIndex int) []string {
	if removeIndex < 0 || removeIndex >= len(guestNames) {
		return guestNames
	}

	oldLen := len(guestNames)

	// TODO: Реализуйте стабильное удаление элемента по индексу removeIndex (порядок остальных должен сохраниться).
	// TODO: Сдвиньте хвост влево с помощью copy (без append и без создания нового массива).
	// TODO: Очистите "освободившийся" последний слот в backing array через clear.
	// TODO: Верните слайс с len на 1 меньше (через реслайс), не создавая новый backing array.

	copy(guestNames[removeIndex:], guestNames[removeIndex+1:])
	clear(guestNames[oldLen-1 : oldLen])
	return guestNames[:oldLen-1]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	guestNames := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &guestNames[i])
	}

	var removeIndex int
	fmt.Fscan(in, &removeIndex)

	guestNames = removeAtStable(guestNames, removeIndex)

	for i, name := range guestNames {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, name)
	}
	fmt.Fprintln(out)
}
