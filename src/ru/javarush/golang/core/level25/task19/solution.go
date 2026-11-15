package main

/*
Вы делаете консольный трекер задач. Внутри доменной логики ошибка должна содержать технический контекст (какая операция
упала), но пользователю это видеть не нужно: ему нужна нормальная, короткая фраза. Поэтому вы заворачиваете “причину”
в OpError, а на границе (main) извлекаете ValidationError через errors.As и строите UX-сообщение не из err.Error(), а
из полей ошибки.

Напишите программу, которая читает из stdin одно слово title. Считайте, что title, равный "-", означает “пустое название”.

Определите ValidationError (value-тип) с полями Field и Msg, а также OpError (pointer-тип) с полями Op string и Err error,
методами Error() string и Unwrap() error.

Реализуйте AddTask(title string) (int, error): если title == "-", вернуть 0 и ошибку
&OpError{Op: "AddTask", Err: ValidationError{Field: "title", Msg: "must not be empty"}}, иначе вернуть id=1 и nil.

На “границе” (в main) при ошибке программа должна печатать две строки:
USER: Название задачи не может быть пустым
DEBUG: <полный err.Error()>

При успехе программа должна печатать:
OK: 1
*/

import (
	"errors"
	"fmt"
)

// ValidationError — value-тип с полями контекста валидации.
type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	// TODO: Верните технический текст ошибки валидации на основе полей Field и Msg.
	// Важно: UX-сообщение будет строиться не из этого текста.
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// OpError — pointer-тип, оборачивает причину и добавляет операцию.
type OpError struct {
	Op  string
	Err error
}

func (e *OpError) Error() string {
	// TODO: Сформируйте полный технический текст ошибки, добавив контекст операции (Op)
	// и текст причины (Err).
	if e == nil {
		return "op error"
	}
	return fmt.Sprintf("%s: %s", e.Op, e.Err)
}

func (e *OpError) Unwrap() error {
	// TODO: Верните вложенную ошибку, чтобы работали errors.As / errors.Is.
	return e.Err
}

func AddTask(title string) (int, error) {
	// TODO: Реализуйте валидацию title и контракт возврата (id, error) по условию задачи.
	// Подсказка по структуре: при ошибке должен возвращаться error с тех. контекстом операции,
	// а внутри — typed error валидации с полями Field/Msg.

	// Временная наивная реализация, не соответствующая требованиям задачи.
	if title == "-" {
		return 0, &OpError{Op: "AddTask", Err: ValidationError{Field: "title", Msg: "must not be empty"}}
	}

	return 1, nil
}

func uxMessage(ve ValidationError) string {
	// TODO: Постройте короткое UX-сообщение пользователю на основе полей ve (Field/Msg),
	// не используя ve.Error().
	if ve.Field == "title" {
		return "Название задачи не может быть пустым"
	}
	return fmt.Sprintf("Некорректный ввод %s: %s", ve.Field, ve.Msg)
}

func main() {
	var title string
	fmt.Scan(&title)

	id, err := AddTask(title)
	if err != nil {
		// Требование проверяющей системы: извлекать typed error валидации через errors.As.
		var ve ValidationError
		if errors.As(err, &ve) {
			fmt.Println("USER:", uxMessage(ve))
		}

		fmt.Println("DEBUG:", err.Error())
		return
	}

	fmt.Println("OK:", id)
}
