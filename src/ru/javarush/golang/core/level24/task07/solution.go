package main

/*
Вы делаете мини‑валидацию логина для консольной регистрации: пользователь вводит логин, а программа либо говорит "OK",
либо объясняет, что не так. Важно сделать это правильно по‑Go: если ошибки нет — возвращать настоящий nil интерфейса error,
а не “хитрый” typed nil.

Создайте тип ошибки ValidationError (поля Field и Msg). Реализуйте метод Error() string на *ValidationError (можно сделать
его nil‑устойчивым, чтобы диагностика была безопасной).
Напишите функцию ValidateLogin(login string) error:
если login пустой или его длина меньше 3, возвращается *ValidationError;
если login корректный, возвращается именно nil как error (не typed nil).

В main вы читаете логин через fmt.Scan, печатаете OK или ERROR: ..., а затем всегда печатаете диагностику err строкой type=<T>, value=<v> через %T и %v.
*/

import "fmt"

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	// TODO: Сформируйте человекочитаемое сообщение об ошибке, используя поля Field и Msg.
	// TODO: Сделайте метод nil-устойчивым, чтобы диагностика через fmt была безопасной даже при typed nil.
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%s : %s", e.Field, e.Msg)
}

func ValidateLogin(login string) error {
	// TODO: Реализуйте валидацию логина по условиям задачи.
	// TODO: При успешной валидации верните именно nil как error (не typed nil).

	if login == "" {
		return &ValidationError{
			Field: "login",
			Msg:   "invalid login",
		}
	}
	if len(login) < 3 {
		return &ValidationError{
			Field: "login",
			Msg:   "invalid login length",
		}
	}
	return nil
}

func main() {
	var login string
	fmt.Scan(&login)

	err := ValidateLogin(login)
	if err == nil {
		fmt.Println("OK")
	} else {
		fmt.Printf("ERROR: %v\n", err)
	}

	fmt.Printf("type=%T, value=%v\n", err, err)
}
