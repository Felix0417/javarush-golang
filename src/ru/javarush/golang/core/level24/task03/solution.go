package main

/*
Вы собираете консольный рендерер текста для небольшого инструмента: иногда текст нужно печатать "как есть", а
иногда — показывать его так, чтобы были видны спецсимволы и пробелы (как будто вы включили режим отладки). Пользователь
вводит режим и строку, а ваша программа выбирает подходящий "форматтер".

Программа читает из stdin две строки: mode и text. mode бывает только "plain" или "quoted". Вы объявляете интерфейс
TextFormatter с методом Format(s string) string, делаете два типа: PlainFormatter (возвращает строку без изменений) и
QuotedFormatter (возвращает строку в кавычках в стиле Go, как %q). Дальше пишете
функцию Render(f TextFormatter, s string) string, которая просто вызывает f.Format(s). В main вы выбираете реализацию по
mode обычным ветвлением и печатаете результат. Если режим неизвестен — печатаете "unknown mode".
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextFormatter interface {
	Format(s string) string
}

type PlainFormatter struct{}

func (PlainFormatter) Format(s string) string { return s }

type QuotedFormatter struct{}

func (QuotedFormatter) Format(s string) string {
	// TODO: Реализуйте форматирование строки в Go-стиле: результат должен быть в кавычках и с экранированием спецсимволов.
	return fmt.Sprintf("%q", s)
}

func Render(f TextFormatter, s string) string {
	// TODO: Реализуйте рендеринг строго через интерфейс: функция должна возвращать результат вызова f.Format(s).
	return f.Format(s)
}

func readLine(r *bufio.Reader) string {
	line, _ := r.ReadString('\n')
	return strings.TrimRight(line, "\r\n")
}

func main() {
	r := bufio.NewReader(os.Stdin)

	mode := readLine(r)
	text := readLine(r)

	var f TextFormatter

	switch mode {
	case "plain":
		f = PlainFormatter{}
	case "quoted":
		f = QuotedFormatter{}
	default:
		fmt.Println("unknown mode")
		return
	}

	fmt.Println(Render(f, text))
}
