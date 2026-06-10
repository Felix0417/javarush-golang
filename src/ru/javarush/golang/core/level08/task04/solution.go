package main

import "fmt"

func main() {
	var fileSize, blockSize int64
	fmt.Scan(&fileSize, &blockSize)

	// TODO: Проверьте входные данные и выведите INVALID, если fileSize < 0 или blockSize <= 0.
	if fileSize < 0 || blockSize <= 0 {
		fmt.Println("INVALID")
		return
	}

	// TODO: Посчитайте количество блоков с округлением вверх, используя только целочисленную арифметику.
	// Текущая версия намеренно не учитывает "хвост" файла.
	blocks := int((fileSize + blockSize - 1) / blockSize)

	// Typed-константа (по заданию должна быть именно типизированной).
	const MaxBlocks uint8 = 255

	// TODO: Сравните blocks с MaxBlocks без ранней конвертации blocks к uint8.
	// Текущая версия намеренно неверная: при больших значениях произойдёт переполнение uint8.
	var blocks8 = int64(blocks)

	if blocks8 > int64(MaxBlocks) {
		fmt.Println("TOO MANY")
		return
	}

	// TODO: Делайте конвертацию к uint8 только в ветке, где blocks гарантированно не превышает MaxBlocks.
	fmt.Println(uint8(blocks8))
}
