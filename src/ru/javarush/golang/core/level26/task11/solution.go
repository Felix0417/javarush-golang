package main

/*
Вы делаете мини-отчёт, который должен одновременно попасть в два места: например, в основной лог и в резервный лог. Писать
две одинаковые строки двумя вызовами — скучно и легко ошибиться, поэтому вы используете io.MultiWriter, чтобы один вызов
записи уходил сразу в два буфера.

Программа читает из stdin число linesCount, а затем linesCount строк (каждая — одно слово без пробелов). Создайте два
независимых bytes.Buffer (например, primaryLog и backupLog), затем общий writer через io.MultiWriter(&primaryLog, &backupLog).
Каждую входную строку записывайте в этот writer отдельной строкой (важен перевод строки).

После этого выведите результат в формате: A: <содержимое первого буфера как есть> B: <содержимое второго буфера как есть>
*/

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var linesCount int
	fmt.Scan(&linesCount)

	// Два независимых буфера.
	var primaryLog, backupLog bytes.Buffer

	// TODO: Создайте общий writer, который пишет одновременно в primaryLog и backupLog через io.MultiWriter.
	w := io.MultiWriter(&primaryLog, &backupLog)

	for i := 0; i < linesCount; i++ {
		var line string
		fmt.Scan(&line)

		// TODO: Записывайте строку в общий writer построчно (с переводом строки) и обязательно проверяйте ошибку записи.
		_, err := fmt.Fprintf(w, "%s\n", line)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Something went wrong in writer %v\n", err)
		}

	}

	// Вывод по формату.
	fmt.Println("A:")
	fmt.Print(primaryLog.String())
	fmt.Println("B:")
	fmt.Print(backupLog.String())
}
