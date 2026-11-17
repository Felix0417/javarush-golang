package main

/*
Представьте, что вы отправляете данные в "капризный" приёмник: он принимает не всё сразу, а только маленькими порциями.
Если ваша функция записи не умеет дописывать хвост, часть данных просто потеряется — и баг будет крайне неприятным.

Сделайте тип LimitedWriter, который специально имитирует такую ситуацию: он хранит записанные байты в памяти и за один
вызов Write принимает не больше Max байт.

Структура должна быть такой: type LimitedWriter struct { Max int; Data []byte }. Метод Write(p []byte) (int, error)
добавляет в Data не больше Max байт за один вызов (если p длиннее, записывается только p[:Max]) и возвращает число реально
записанных байт.

Затем реализуйте функцию WriteAll(w io.Writer, p []byte) error, которая гарантированно записывает в w вообще все байты p,
делая столько вызовов Write, сколько понадобится.

В main() прочитайте одно слово из stdin (например, message), добавьте к нему перевод строки, запишите через WriteAll
в LimitedWriter с Max = 3, а затем выведите накопленные данные LimitedWriter в stdout как строку.
*/

import (
	"fmt"
	"io"
	"os"
)

type LimitedWriter struct {
	Max  int
	Data []byte
}

// Write принимает за раз не больше Max байт и накапливает их в памяти.
func (w *LimitedWriter) Write(p []byte) (int, error) {
	// TODO: Реализуйте ограничение "не больше Max байт за один вызов Write".
	// TODO: Реализуйте накопление записанных данных в w.Data через append, не затирая ранее записанное.

	n := len(p)
	if w.Max >= 0 && n > w.Max {
		n = w.Max
	}

	// Намеренно неполная реализация: перезаписывает данные вместо накопления.
	w.Data = append(w.Data, p[:n]...)
	return n, nil
}

// WriteAll гарантированно дозаписывает весь p, учитывая частичные записи.
func WriteAll(w io.Writer, p []byte) error {
	// TODO: Реализуйте запись всего p через несколько вызовов w.Write (если нужно).
	// TODO: Учитывайте (n, err): при err != nil возвращайте ошибку, иначе продолжайте с p[n:].
	// TODO: Добавьте защиту от бесконечного цикла: если n == 0 и err == nil — верните ошибку.
	for len(p) > 0 {
		bytes, err := w.Write(p)
		if err != nil {
			return err
		} else if bytes == 0 {
			return fmt.Errorf("Wrote 0 bytes")
		}
		p = p[bytes:]
	}
	return nil
}

func main() {
	var word string
	fmt.Fscan(os.Stdin, &word)

	lw := &LimitedWriter{Max: 3}
	_ = WriteAll(lw, []byte(word+"\n"))

	fmt.Print(string(lw.Data))
}
