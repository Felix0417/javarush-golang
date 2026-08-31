package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)

	list := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&list[i])
	}

	// Используем пакет slices, но ключевая логика редактирования списка вынесена в TODO ниже.
	list = slices.Clone(list)

	var q int
	fmt.Scan(&q)

	for k := 0; k < q; k++ {
		var cmd string
		fmt.Scan(&cmd)

		if cmd == "DEL" {
			var i, j int
			fmt.Scan(&i, &j)

			// TODO: Реализуйте удаление диапазона из list по команде DEL:
			// TODO: 1) проверьте корректность границ (чтобы не было panic),
			// TODO: 2) удалите элементы полуинтервала, обновив list результатом операции из пакета slices,
			// TODO: 3) некорректную команду полностью игнорируйте (list не меняется).
			if i < 0 || j < 0 || i > j || j > len(list) {
				continue
			}
			list = slices.Delete(list, i, j)
		}

		if cmd == "INS" {
			var i int
			var x string
			fmt.Scan(&i, &x)

			// TODO: Реализуйте вставку слова x в позицию i по команде INS:
			// TODO: 1) проверьте корректность индекса (чтобы не было panic),
			// TODO: 2) вставьте элемент, обновив list результатом операции из пакета slices,
			// TODO: 3) некорректную команду полностью игнорируйте (list не меняется).
			// TODO: Важно: не создавайте новую переменную list через :=, переиспользуйте существующую.
			if i < 0 || i > len(list) {
				continue
			}
			list = slices.Insert(list, i, x)
		}

		// Неизвестная команда — игнорируем.
	}

	for i := 0; i < len(list); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(list[i])
	}
	fmt.Println()
}
