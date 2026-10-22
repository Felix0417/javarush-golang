package apperr

import (
	"errors"
)

// Sentinel-ошибки — по ним классифицируем через errors.Is.
var (
	ErrValidation = errors.New("validation")
	ErrNotFound   = errors.New("not found")
	ErrIO         = errors.New("io")
	ErrInternal   = errors.New("internal")
)

// UserMessage возвращает короткое сообщение для пользователя.
// Важно: классификация должна работать для wrap/join.
func UserMessage(err error) string {
	if err == nil {
		return "OK"
	}

	// TODO: Реализуйте классификацию ошибки через errors.Is (включая случаи wrap и join) и верните корректное сообщение для пользователя.
	if errors.Is(err, ErrValidation) {
		return "Неверный ввод"
	}

	if errors.Is(err, ErrNotFound) {
		return "Не найдено"
	}

	if errors.Is(err, ErrIO) {
		return "Ошибка ввода/вывода"
	}

	if errors.Is(err, ErrInternal) {
		return "Внутренняя ошибка"
	}

	return "Внутренняя ошибка"
}
