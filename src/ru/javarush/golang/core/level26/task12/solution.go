package main

/*
Вы собираете маленький I/O-конвейер, как в “настоящих” пайплайнах: сначала ограничиваем поток первыми байтами, потом
“подслушиваем” его в два журнала, а дальше ваш код вычитывает поток чанками и показывает пользователю то, что получилось.
Весь фокус — в правильном соединении io.LimitReader, io.TeeReader и io.MultiWriter.

Программа читает из stdin строку sourceText (без пробелов), целое число byteLimit (лимит байт) и целое число chunkSize
(размер чанка чтения, chunkSize > 0).

Постройте конвейер:
- сначала strings.NewReader(sourceText),
- затем io.LimitReader(..., int64(byteLimit)),
- затем io.TeeReader(...), который пишет всё реально прочитанное одновременно в два буфера через io.MultiWriter(&log1, &log2).

Дальше полностью вычитайте получившийся reader в цикле (чанками), соберите “показанный пользователю” текст в отдельный
bytes.Buffer и выведите:
- shown=<через %q>
- log1=<через %q>
- log2=<через %q>
- bytes=<количество байт в shown>
*/

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var sourceText string
	var byteLimit, chunkSize int
	if _, err := fmt.Fscan(os.Stdin, &sourceText, &byteLimit, &chunkSize); err != nil {
		fmt.Printf("error=%v\n", err)
		return
	}

	if chunkSize <= 0 {
		fmt.Printf("error=%v\n", "chunkSize must be > 0")
		return
	}
	if byteLimit < 0 {
		fmt.Printf("error=%v\n", "byteLimit must be >= 0")
		return
	}

	var log1, log2 bytes.Buffer
	var shown bytes.Buffer

	// TODO: Соберите I/O-конвейер в требуемом порядке через strings.NewReader, io.LimitReader и io.TeeReader.
	// TODO: В качестве writer для TeeReader используйте io.MultiWriter(&log1, &log2).
	r := strings.NewReader(sourceText)
	limitR := io.LimitReader(r, int64(byteLimit))
	w := io.MultiWriter(&log1, &log2)

	teeR := io.TeeReader(limitR, w)

	// TODO: Читайте данные только через Read в цикле, используя буфер размера chunkSize.
	// TODO: Не используйте io.Copy/io.ReadAll.
	buf := make([]byte, chunkSize)
	for {
		k, err := teeR.Read(buf)

		if k > 0 {
			// TODO: Добавляйте в shown только реально прочитанные байты buf[:k].
			shown.Write(buf[:k])
		}

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error=%v\n", err)
			return
		}

		if k == 0 {
			// Защита от потенциального зависания при некорректной реализации reader'а.
			fmt.Printf("error=%v\n", "unexpected empty read")
			return
		}
	}

	fmt.Printf("shown=%q\n", shown.String())
	fmt.Printf("log1=%q\n", log1.String())
	fmt.Printf("log2=%q\n", log2.String())
	fmt.Printf("bytes=%d\n", shown.Len())
}
