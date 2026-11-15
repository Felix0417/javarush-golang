package main

/*
Вы пишете функцию поиска по ID, как в настоящем сервисе: при “не найдено” возвращается typed error, причём метод Error()
реализован на указателе (*NotFoundError). На границе (в main) вы хотите достать именно *NotFoundError через errors.As, и
тут появляется важный нюанс: целевая переменная сама должна быть указателем, а в errors.As передаётся адрес этой
переменной (получается “указатель на указатель”).

Напишите программу, которая читает из stdin одно целое число requestedID.

Определите typed error NotFoundError с полем ID int и методом Error() string, реализованным на *NotFoundError
(receiver — указатель). Реализуйте функцию FindByID(id int) error: если id равен 42, функция возвращает nil, иначе
возвращает ошибку, обёрнутую как fmt.Errorf("FindByID: %w", &NotFoundError{ID: id}).

В main программа вызывает FindByID. Если err == nil, вывод должен быть FOUND. Если err != nil, программа должна
извлечь *NotFoundError через errors.As и вывести NOT_FOUND <id>.
*/

import (
	"errors"
	"fmt"
)

// NotFoundError — typed error с контекстом (ID).
type NotFoundError struct {
	ID int
}

// Error() реализован на *NotFoundError (receiver — указатель).
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("id %d not found", e.ID)
}

// FindByID ищет объект по ID.
func FindByID(id int) error {
	// TODO: Реализуйте логику поиска:
	// TODO: 1) при id == 42 верните nil
	// TODO: 2) иначе верните typed error *NotFoundError{ID: id}, обёрнутую через fmt.Errorf с %w и префиксом "FindByID:"
	if id == 42 {
		return nil
	}

	return fmt.Errorf("FindByID: %w", &NotFoundError{ID: id})
}

func main() {
	var requestedID int
	fmt.Scan(&requestedID)

	err := FindByID(requestedID)
	if err == nil {
		fmt.Print("FOUND")
		return
	}

	var nf *NotFoundError

	// TODO: Извлеките из err именно *NotFoundError через errors.As.
	// TODO: Обратите внимание, что целевая переменная должна быть указателем, и в errors.As передаётся адрес этой переменной.
	// TODO: После извлечения выведите NOT_FOUND <id>, где <id> берётся из nf.ID.
	if errors.As(err, &nf) {
		fmt.Printf("NOT_FOUND %d\n", nf.ID)
	}
}
