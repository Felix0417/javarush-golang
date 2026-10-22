package app

import "errors"

// Sentinel-ошибки приложения: по ним делаем классификацию через errors.Is.
var (
	ErrValidation = errors.New("validation")
	ErrNotFound   = errors.New("not found")
	ErrIO         = errors.New("io")
	ErrInternal   = errors.New("internal")
)

// Class возвращает класс ошибки по причине в цепочке (wrapping/join), а не по тексту.
func Class(err error) string {
	if err == nil {
		return "ok"
	}

	// TODO: Реализуйте классификацию только через errors.Is для всех sentinel-ошибок:
	// TODO: ErrValidation, ErrNotFound, ErrIO, ErrInternal.
	// TODO: Должны корректно определяться случаи wrapping через %w и errors.Join.
	// TODO: Нельзя определять класс по тексту ошибки.

	switch {
	case errors.Is(err, ErrValidation):
		return "validation"
	case errors.Is(err, ErrNotFound):
		return "notfound"
	case errors.Is(err, ErrIO):
		return "io"
	case errors.Is(err, ErrInternal):
		return "internal"
	default:
		return "internal"
	}
}
