package main

import (
	"bufio"
	"fmt"
	"os"
)

func removeAtStable(contacts []string, removeIndex int) []string {
	// TODO: Реализуйте стабильное удаление элемента по индексу removeIndex (с сохранением порядка).
	// TODO: Сдвиг элементов сделайте через copy.
	// TODO: Перед уменьшением длины очистите освободившийся слот в backing array через clear.
	// TODO: Верните результат через reslice (уменьшение длины на 1).
	copy(contacts[removeIndex:], contacts[removeIndex+1:])
	clear(contacts[len(contacts)-1:])
	return contacts[:len(contacts)-1]
}

func removeAtUnstable(contacts []string, removeIndex int) []string {
	// TODO: Реализуйте быстрое (unstable) удаление элемента по индексу removeIndex (порядок может измениться).
	// TODO: Используйте подход "swap-with-last": положите последний элемент на место удаляемого.
	// TODO: Перед уменьшением длины очистите освобождённый слот в backing array через clear.
	// TODO: Верните результат через reslice (уменьшение длины на 1).
	contacts[removeIndex] = contacts[len(contacts)-1]
	clear(contacts[len(contacts)-1:])
	return contacts[:len(contacts)-1]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	contacts := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &contacts[i])
	}

	var mode string
	fmt.Fscan(in, &mode)

	var k int
	fmt.Fscan(in, &k)

	// Перевод "человеческого" номера в индекс — ровно в одном месте
	removeIndex := k - 1

	res := contacts
	if removeIndex >= 0 && removeIndex < len(contacts) {
		if mode == "stable" {
			res = removeAtStable(res, removeIndex)
		} else if mode == "unstable" {
			res = removeAtUnstable(res, removeIndex)
		}
	}

	for _, c := range res {
		fmt.Fprintln(out, c)
	}
}
