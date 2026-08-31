package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)

	// Читаем строку целиком до '\n' (если '\n' нет — возможен io.EOF, но строка может быть полезной).
	s, err := rd.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		return
	}

	found := false

	// Важно: range по строке идёт по рунам, а i — байтовый индекс.
	for i, r := range s {
		// TODO: Найдите управляющие символы '\t', '\r', '\n' при обходе строки по рунам.
		// TODO: Для каждого найденного символа выведите отдельную строку в требуемом формате и отметьте found=true.
		switch r {
		case '\t', '\r', '\n':
			found = true
			fmt.Printf("pos=%d rune=%q code=%d\n", i, r, r)
		}
	}

	if !found {
		fmt.Println("clean")
	}
}
