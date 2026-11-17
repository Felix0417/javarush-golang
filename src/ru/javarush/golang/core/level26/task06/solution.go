package main

/*
Вы пишете маленький генератор "карточки места" для путеводителя: на вход приходит город и страна, а на выход — аккуратный
двухстрочный блок текста. По условию вы сначала формируете весь текст в памяти, и только потом печатаете его целиком — как
будто готовите сообщение для отправки.

Программа читает из stdin два слова без пробелов: cityName и countryName. Затем формирует многострочный текст через bytes.Buffer,
направляя fmt.Fprintln/Fprintf в буфер. Формат должен быть строго таким:
City: <city>
Country: <country>

После формирования выведите содержимое буфера в stdout одним вызовом fmt.Print или fmt.Println.
*/

import (
	"bytes"
	"fmt"
)

func main() {
	var cityName, countryName string
	fmt.Scan(&cityName, &countryName)

	var buf bytes.Buffer // нулевое значение, без make/new

	// TODO: Сформируйте в buf ровно две строки в требуемом формате (с переводом строки после каждой строки).

	fmt.Fprintf(&buf, "City: %s\n", cityName)
	fmt.Fprintf(&buf, "Country: %s\n", countryName)

	// Один финальный вывод после полной сборки результата в памяти.
	fmt.Println(buf.String())
}
