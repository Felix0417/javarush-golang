package main

import (
	"bufio"
	f "fmt"
	"os"
	"strings"
)

func main() {
	fmt := "hello," // по условию нельзя переименовывать

	in := bufio.NewReader(os.Stdin)
	name, _ := in.ReadString('\n')

	name = strings.TrimSpace(name)

	// TODO: Нормализуйте имя пользователя через strings.TrimSpace (убрать пробелы и перевод строки).
	// TODO: Импортируйте пакет "fmt" с алиасом, чтобы не конфликтовать с переменной fmt, и печатайте через алиас.
	// TODO: Выведите ровно одну строку в формате: hello, <name>.
	f.Println(fmt, name)
}
