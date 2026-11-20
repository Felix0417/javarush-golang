package main

/*
Вы пишете консольную утилиту, которая показывает “превью” содержимого: например, первые N байт файла или сообщения.
Чтобы не читать лишнего, вы ограничиваете поток с помощью io.LimitReader, а затем вычитываете уже ограниченный поток
кусочками в буфер и собираете результат.

Программа читает из stdin два значения: строку sourceText (без пробелов) и целое число previewLimit. Затем создаёт
источник через strings.NewReader(sourceText), оборачивает его в io.LimitReader(..., int64(previewLimit)), полностью
вычитывает получившийся reader в цикле через Read (чанками фиксированного размера) и накапливает показанный текст
в bytes.Buffer.

В конце выведите две строки: preview=<данные в кавычках через %q> bytes=<количество реально прочитанных байт>
*/

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var sourceText string
	var previewLimit int

	fmt.Scan(&sourceText, &previewLimit)

	sourceReader := strings.NewReader(sourceText)
	limited := io.LimitReader(sourceReader, int64(previewLimit))

	var (
		out       bytes.Buffer
		bytesRead int
	)

	// TODO: Реализовать полное вычитывание limited в цикле через Read
	// TODO: Использовать заранее выделенный буфер фиксированного размера и накапливать результат по buf[:k]
	// TODO: Корректно обрабатывать io.EOF как нормальное завершение и любые другие ошибки как ошибку
	// TODO: Корректно посчитать bytesRead как число реально прочитанных байт (не больше previewLimit)
	var buf = make([]byte, 8)

	for {
		k, err := limited.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("error=%v\n", err)
			return
		} else if err == io.EOF {
			break
		}

		if k > 0 {
			out.Write(buf[:k])
			bytesRead += k
		}
	}

	fmt.Printf("preview=%q\n", out.String())
	fmt.Printf("bytes=%d\n", bytesRead)
}
