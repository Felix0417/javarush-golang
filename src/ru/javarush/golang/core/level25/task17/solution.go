package main

/*
Вы делаете проверку формы ввода в консольном сервисе. Иногда ошибка “типизированная” и содержит поля (например, какое поле
формы сломалось), а иногда — обычная текстовая. При этом ошибки могут быть обёрнуты “по дороге”, и на границе программы
вы хотите надёжно определить: есть ли внутри цепочки именно ValidationError, и если да — достать из него поля, а не гадать
по тексту.

Напишите программу, которая читает из stdin режим работы.
- Если введено typed, программа читает ещё два слова: fieldName и messageText. Затем создаёт ошибку ValidationError с этими
значениями, оборачивает её через fmt.Errorf("wrap: %w", err) и пытается извлечь ValidationError с помощью errors.As.
- Если введено plain, программа читает ещё одно слово messageText, создаёт обычную ошибку через errors.New(messageText),
оборачивает её через fmt.Errorf("wrap: %w", err) и также пытается извлечь ValidationError.

Вывод должен быть стабильным: если ValidationError найден, программа должна вывести строку FIELD=<field> MSG=<msg>.
Если ValidationError не найден, программа должна вывести NOT_VALIDATION.
*/

import (
	"errors"
	"fmt"
	"os"
)

// ValidationError — типизированная ошибка валидации с полями контекста.
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	// Текст не используется для распознавания; важен сам тип + поля.
	return e.Field + ": " + e.Msg
}

func main() {
	var mode string
	fmt.Fscan(os.Stdin, &mode)

	var err error
	if mode == "typed" {
		var fieldName, messageText string
		fmt.Fscan(os.Stdin, &fieldName, &messageText)

		err = &ValidationError{
			Field: fieldName,
			Msg:   messageText,
		}
	} else { // plain
		var messageText string
		fmt.Fscan(os.Stdin, &messageText)

		err = errors.New(messageText)
	}

	wrapped := fmt.Errorf("wrap: %w", err)

	var target *ValidationError

	found := false
	// TODO: С помощью errors.As проверьте, есть ли внутри wrapped ошибка типа *ValidationError.
	// TODO: Если да — запишите найденную ошибку в target и установите found = true.
	if errors.As(wrapped, &target) {
		found = true
	}

	if found {
		fmt.Printf("FIELD=%s MSG=%s\n", target.Field, target.Msg)
		return
	}

	fmt.Print("NOT_VALIDATION\n")
}
