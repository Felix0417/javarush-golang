package main

/*
Представьте, что вы пишете маленький модуль для чат-бота: бот не обязан всегда говорить в консоль — иногда он пишет в файл,
иногда в сеть, иногда в память для тестов. Поэтому приветствие нужно уметь отправлять в любой io.Writer.

Реализуйте функцию Greet(w io.Writer, name string) error, которая записывает в w строку приветствия строго в формате
"Hello, <name>!" и добавляет перевод строки в конце.

В main() ваша программа читает из stdin одно слово — имя пользователя (например, userName) — и вызывает
Greet(os.Stdout, userName). Если Greet вернула ошибку, её нужно аккуратно вывести в stderr и завершить программу без паники.
*/

import (
	"fmt"
	"io"
	"os"
)

// Greet пишет приветствие в любой io.Writer и возвращает ошибку записи.
func Greet(w io.Writer, name string) error {
	// TODO: Запишите приветствие строго в формате из условия, используя fmt.Fprintf или fmt.Fprintln (запись должна идти именно в w). Не забудьте перевод строки в конце. Верните ошибку записи (если она возникла).
	_, err := fmt.Fprintf(w, "Hello, %s!\n", name)
	return err
}

func main() {
	var userName string
	fmt.Fscan(os.Stdin, &userName)

	// os.Stdout передаём как io.Writer, чтобы не привязываться к консоли.
	if err := Greet(os.Stdout, userName); err != nil {
		// Ошибку печатаем в stderr и выходим без паники.
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
