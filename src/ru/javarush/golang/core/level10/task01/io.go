package main

import (
	"fmt"
)

// readTwoInts отвечает только за ввод: читает ровно два int и возвращает их.
func readTwoInts() (int, int) {
	// TODO: Прочитайте из stdin ровно два целых числа (int) и верните их.
	var first, second int
	_, err := fmt.Scan(&first, &second)
	if err != nil {
		return 0, 0
	}
	return first, second
}
