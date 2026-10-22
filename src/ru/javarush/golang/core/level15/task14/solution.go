package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	priceByCode := make(map[string]int, n)

	for i := 0; i < n; i++ {
		var code string
		var price int
		fmt.Fscan(in, &code, &price)

		// Повторяющийся код перезаписываем: важна последняя цена.
		priceByCode[code] = price
	}

	// TODO: Сделайте стабильный вывод: соберите все ключи map в отдельный слайс,
	// TODO: отсортируйте ключи по возрастанию и напечатайте пары (key value)
	// TODO: строго в порядке отсортированных ключей (не через range по map).

	tempSlice := make([]string, 0, len(priceByCode))
	for k := range priceByCode {
		tempSlice = append(tempSlice, k)
	}

	sort.Strings(tempSlice)

	for _, code := range tempSlice {
		fmt.Fprintln(out, code, priceByCode[code])
	}
}
