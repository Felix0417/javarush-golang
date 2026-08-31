package main

import (
	"fmt"
	"strings"
)

func main() {
	raw := " \t  Go is fun \n  "

	// TODO: Получите строку clean из raw с помощью strings.TrimSpace(raw).
	// Сейчас обрезаем только пробелы ' ' по краям — этого недостаточно (таб и перевод строки останутся).
	clean := strings.TrimSpace(raw)

	// TODO: Выведите raw и clean так, чтобы были видны управляющие символы (используйте формат %q).
	fmt.Printf("%q\n", raw)
	fmt.Printf("%q\n", clean)

	// TODO: Выведите длины raw и clean в байтах, по одному числу на строку.
	fmt.Printf("%d\n", len(raw))
	fmt.Printf("%d\n", len(clean))
}
