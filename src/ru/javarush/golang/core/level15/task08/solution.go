package main

import "fmt"

func main() {
	notes := make(map[string]string) // хранилище заметок key -> value

	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "set":
			var key, value string
			fmt.Scan(&key, &value)

			// TODO: Реализуйте сохранение заметки в map notes по ключу key.
			// TODO: Если value равно "-", сохраните пустую строку "" (это валидное значение).
			// TODO: Убедитесь, что эта команда печатает ровно одну строку результата (пустую строку).
			if value == "-" {
				notes[key] = ""
			} else {
				notes[key] = value
			}
			fmt.Println(notes[key])
		case "get":
			var key string
			fmt.Scan(&key)

			// TODO: Реализуйте получение значения ТОЛЬКО через ok-идиому: v, ok := notes[key].
			// TODO: Если ключа нет — печатайте "no such key".
			// TODO: Если ключ есть — печатайте: value: "<строка>" (кавычки обязательны).
			// TODO: Сделайте так, чтобы пустая строка была видна в выводе (например, через форматирование %q).
			v, ok := notes[key]
			if !ok {
				fmt.Println("no such key")
			} else {
				fmt.Printf("value: %q\n", v)
			}
		case "exists":
			var key string
			fmt.Scan(&key)

			// TODO: Реализуйте проверку существования ключа через ok-идиому: _, ok := notes[key].
			// TODO: Печатайте ровно true или false в зависимости от наличия ключа.
			_, ok := notes[key]
			fmt.Println(ok)

		case "del":
			var key string
			fmt.Scan(&key)

			// TODO: Реализуйте удаление: сначала проверьте существование ключа через ok-идиому.
			// TODO: Если ключ существовал — удалите его через delete(notes, key) и выведите "deleted".
			// TODO: Если ключа не было — выведите "missing".
			if _, ok := notes[key]; ok {
				delete(notes, key)
				fmt.Println("deleted")
			} else {
				fmt.Println("missing")
			}
		}
	}
}
