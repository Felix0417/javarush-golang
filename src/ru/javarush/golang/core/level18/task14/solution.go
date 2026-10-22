package main

import (
	"fmt"
)

// run — вся "бизнес-логика": читаем x и делим 100 / x без проверок на 0.
func run() int {
	var x int
	fmt.Scan(&x)

	// TODO: Выполните деление строго как 100 / x (int) без каких-либо if-проверок на x == 0.
	// Важно: деление должно быть именно аварийным при x == 0 (runtime panic), чтобы safeRun мог это перехватить.
	return 100 / x
}

// safeRun — граница приложения: превращаем panic в обычную ошибку.
func safeRun() (result int, err error) {
	// TODO: Реализуйте "границу приложения":
	// 1) поставьте defer с recover()
	// 2) при панике установите err (через именованные возвращаемые значения)
	// 3) не допускайте проброса паники наружу
	// Подсказка: сигнатуру (result int, err error) менять нельзя.
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	result = run()
	return result, nil
}

func main() {
	result, err := safeRun()
	if err != nil {
		fmt.Printf("app error: %v\n", err)
		return
	}

	fmt.Printf("result: %d\n", result)
}
