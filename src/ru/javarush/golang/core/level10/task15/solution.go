package main

import (
	"errors"
	"fmt"
	"os"
)

// init() намеренно не используем: инициализация должна быть явной.

func initApp(appTag string) (string, error) {
	// TODO: Реализуйте явную инициализацию приложения.
	// TODO: Проверьте корректность appTag (пустая строка должна приводить к ошибке с нужным текстом).
	// TODO: При успехе верните строку-префикс, из которой main сможет собрать итоговый вывод.
	if appTag == "" {
		return "", errors.New("empty app tag")
	}
	return appTag, nil
}

func main() {
	var appTag string
	fmt.Fscan(os.Stdin, &appTag) // читаем ровно одно слово

	prefix, err := initApp(appTag)
	if err != nil {
		fmt.Print("init error: ", err)
		return
	}

	fmt.Print(prefix, "ready")
}
