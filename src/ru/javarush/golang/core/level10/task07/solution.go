package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	var method string
	var code int

	fmt.Scan(&method, &code)

	// TODO: Нормализуйте HTTP-метод через strings.ToUpper.
	// В шаблоне метод намеренно оставлен без нормализации.
	method = strings.ToUpper(method)

	// TODO: Получите текст статуса через http.StatusText(code).
	// TODO: Если http.StatusText(code) вернул пустую строку — используйте "UNKNOWN".
	// В шаблоне намеренно используется фиксированное значение, чтобы задача не была решена.
	status := http.StatusText(code)

	if status == "" {
		status = "UNKNOWN"
	}

	fmt.Printf("%s %s\n", method, status)

	// code пока не используется в шаблоне.
}
