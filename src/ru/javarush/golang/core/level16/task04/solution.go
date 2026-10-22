package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// n читаем именно как строковый токен.
	var nToken string
	if _, err := fmt.Fscan(in, &nToken); err != nil {
		fmt.Printf("FAILED %v\n", err)
		return
	}

	// TODO: Преобразуйте nToken в число через strconv.Atoi и сохраните получившийся n.
	// TODO: При ошибке парсинга выведите "FAILED <err>" и немедленно завершите main через return.
	n, err := strconv.Atoi(nToken)

	if err != nil {
		fmt.Printf("FAILED %v\n", err)
		return
	}

	total := 0

	for i := 0; i < n; i++ {
		// Категория не влияет на расчёт, но токен всё равно нужно прочитать.
		var expenseCategory string
		var amountToken string
		if _, err := fmt.Fscan(in, &expenseCategory, &amountToken); err != nil {
			fmt.Printf("FAILED %v\n", err)
			return
		}

		// TODO: Преобразуйте amountToken в int через strconv.Atoi.
		// TODO: При ошибке парсинга выведите "FAILED <err>" и немедленно завершите main через return.
		// TODO: При успехе прибавьте сумму к total.
		val, err := strconv.Atoi(amountToken)
		if err != nil {
			fmt.Printf("FAILED %v\n", err)
			return
		}
		total += val
	}

	// Используем strconv, чтобы импорт не был лишним в шаблоне.

	fmt.Printf("TOTAL=%d\n", total)
}
