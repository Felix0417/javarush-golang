package main

/*
Представьте, что вы печатаете одну строку чека в магазине. У строки фиксированный формат, и если в ней будет лишний
пробел или потеряется перевод строки — тесты кассы “завалят” сборку. Поэтому вы делаете функцию, которая пишет строку в
любой io.Writer, а затем в main() проверяете, что получилось ровно то, что ожидали.

Реализуйте функцию WriteLineItem(w io.Writer, title string, qty int, price int) error. Она должна записывать в w одну
строку формата: ITEM=<title> QTY=<qty> PRICE=<price> TOTAL=<qty*price> и добавлять перевод строки в конце. Ошибка записи
должна возвращаться наружу.

В main() вызовите WriteLineItem, передав bytes.Buffer как место вывода. Затем сравните получившуюся строку got с ожидаемой
строкой want. Если строки совпали — выведите OK. Если нет — выведите MISMATCH, затем got и want (обе строки печатайте
через %q, чтобы был виден \n).
*/

import (
	"bytes"
	"fmt"
	"io"
)

// WriteLineItem пишет одну строку чека в writer.
func WriteLineItem(w io.Writer, title string, qty int, price int) error {
	// TODO: Записать в w ровно одну строку в требуемом формате и обязательно добавить '\n' в конце. Ошибку записи вернуть наружу.
	_, err := fmt.Fprintf(w, "ITEM=%s QTY=%d PRICE=%d TOTAL=%d\n", title, qty, price, qty*price)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var buf bytes.Buffer

	// Пишем в буфер (не в stdout) — так удобно проверять точное совпадение строки.
	if err := WriteLineItem(&buf, "MILK", 3, 70); err != nil {
		fmt.Printf("write error: %v\n", err)
		return
	}

	got := buf.String()
	want := "ITEM=MILK QTY=3 PRICE=70 TOTAL=210\n"

	if got == want {
		fmt.Println("OK")
		return
	}

	fmt.Println("MISMATCH")
	fmt.Printf("%q\n", got)
	fmt.Printf("%q\n", want)
}
