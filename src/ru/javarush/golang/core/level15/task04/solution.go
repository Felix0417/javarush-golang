package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q)

	// По условию карта должна начинаться как nil-map.
	var book map[string]string

	for i := 0; i < q; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		switch cmd {
		case "set":
			var name, phone string
			fmt.Fscan(in, &name, &phone)

			// TODO: Реализуйте команду set: при первом set инициализируйте book через make (ленивая инициализация),
			// TODO: затем сохраните/обновите phone по ключу name.
			if book == nil {
				book = make(map[string]string)
			}
			book[name] = phone

		case "get":
			var name string
			fmt.Fscan(in, &name)

			// TODO: Реализуйте команду get: получите значение через book[name],
			// TODO: если результат равен пустой строке "", выведите MISSING, иначе выведите телефон.

			result := book[name]
			if result == "" {
				fmt.Fprintln(out, "MISSING")
			} else {
				fmt.Fprintln(out, result)
			}

		case "del":
			var name string
			fmt.Fscan(in, &name)

			// TODO: Реализуйте команду del: удалите запись через delete(book, name) без дополнительных проверок.
			delete(book, name)

		case "len":
			// TODO: Реализуйте команду len: выведите текущее количество записей как len(book)
			// TODO: (в том числе для nil-map должно получаться 0).

			fmt.Fprintln(out, len(book))
		}
	}
}
