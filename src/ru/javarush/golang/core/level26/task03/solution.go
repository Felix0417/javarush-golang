package main

/*
Вы делаете “санитайзер” для текстового потока: на вход приходит текст (например, из файла или stdin), а на выход нужно
отдать тот же поток, но без обычных пробелов ' ' — чтобы потом сравнивать строки, считать контрольные суммы или отправлять
данные в систему, где пробелы запрещены.

Реализуйте функцию CopyWithoutSpaces(r io.Reader, w io.Writer) error. Она должна читать байты из r небольшими порциями
в буфер, выкидывать из прочитанного только пробелы ' ' (ASCII 0x20), а остальное записывать в w. Учтите важную
особенность: Write может записать меньше, чем вы ему передали, поэтому подготовленный кусок нужно “дозаписывать” в цикле,
пока он не уйдёт целиком.

В main() вызовите CopyWithoutSpaces(os.Stdin, os.Stdout). Если чтение или запись завершились ошибкой — выведите ошибку в stderr.
*/

import (
	"fmt"
	"io"
	"os"
)

func CopyWithoutSpaces(r io.Reader, w io.Writer) error {
	buf := make([]byte, 32*1024)

	for {
		n, err := r.Read(buf)

		if n > 0 {
			chunk := buf[:n]
			clean := make([]byte, 0, len(chunk))
			// TODO: Удалите из chunk только обычные пробелы ' ' (ASCII 0x20), сохранив порядок остальных байт.
			// TODO: Запишите подготовленные данные в w, учитывая, что Write может записать меньше, чем передано (partial write).
			for _, v := range chunk {
				if v != ' ' {
					clean = append(clean, v)
				}

			}

			for len(clean) > 0 {
				bytes, werr := w.Write(clean)
				if werr != nil {
					return werr
				}
				clean = clean[bytes:]
			}

		}

		// TODO: Корректно обработайте сценарий, когда err != nil и при этом n > 0:
		// TODO: сначала нужно обработать/записать эти n байт, а затем вернуть ошибку чтения.

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

func main() {
	if err := CopyWithoutSpaces(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
