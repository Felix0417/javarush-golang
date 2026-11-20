package main

/*
Представьте, что вы отлаживаете чтение данных: ваш код читает поток, а параллельно хочется иметь “чёрный ящик” — копию
всего, что реально было прочитано. Для этого идеально подходит io.TeeReader: он отдаёт байты вашему коду и одновременно
пишет их в другой io.Writer.

Программа читает из stdin строку payload (без пробелов) и целое число chunkSize (размер чанка чтения, chunkSize > 0).
Создайте strings.NewReader(payload) и оберните его в io.TeeReader, который пишет всё прочитанное в bytes.Buffer (например, debugCopy).
Затем полностью вычитайте поток в цикле через Read чанками размером chunkSize, собирая результат чтения в отдельный
буфер (например, actuallyRead).

После окончания чтения выведите: read=<всё прочитанное через %q> copy=<содержимое отладочной копии через %q>
*/

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var payload string
	var chunkSize int
	fmt.Scan(&payload, &chunkSize)

	src := strings.NewReader(payload)

	var debugCopy bytes.Buffer
	tee := io.TeeReader(src, &debugCopy)

	var actuallyRead bytes.Buffer
	tmp := make([]byte, chunkSize)

	for {
		n, err := tee.Read(tmp)

		if n > 0 {
			actuallyRead.Write(tmp[:n])
		}

		// TODO: Реализуйте корректный цикл чтения чанками до конца потока.
		// TODO: Корректно обработайте io.EOF как штатное завершение чтения.
		// TODO: При любой другой ошибке чтения выведите понятное сообщение и завершите программу.
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("something went wrong in reader %v\n", err)
			return
		}
	}

	fmt.Printf("read=%q\n", actuallyRead.String())
	fmt.Printf("copy=%q\n", debugCopy.String())
}
