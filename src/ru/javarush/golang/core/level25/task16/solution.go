package main

/*
У вас есть мини-"база пользователей" в map[int]string. Пользователь вводит ID как строку, и дальше возможны два типа
проблем: строка вообще не число (ошибка парсинга) или число корректное, но такого пользователя нет (ошибка "не найдено").
Вы хотите сохранить эти причины как typed errors, а сверху добавить контекст операции через OpError, не теряя первопричину.

Напишите программу, которая хранит "базу пользователей" в map[int]string (например: 1 → "Ann", 2 → "Bob"). Программа
читает из stdin один токен rawID (строку) и пытается получить имя пользователя по этому ID.

Реализуйте ParseID(raw string) (int, error): при ошибке парсинга возвращает typed error InvalidIDError с полем Input.
Реализуйте FindUser(users map[int]string, id int) (string, error): если ключа нет, возвращает typed error UserNotFoundError
с полем ID. Реализуйте GetUserName(users map[int]string, rawID string) (string, error): вызывает ParseID и FindUser и
оборачивает ошибки в OpError (формат Op + Err), при этом OpError должен поддерживать Unwrap().

В main при InvalidIDError выводится INVALID ID: <Input>, при UserNotFoundError выводится NOT FOUND: <ID>, при успехе
выводится NAME: <имя>.
*/

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// InvalidIDError — typed error для ошибки парсинга (важно: поле доступно коду, не через текст).
type InvalidIDError struct {
	Input string
}

func (e InvalidIDError) Error() string {
	// Текст ошибки не должен использоваться для принятия решений в программе.
	return "invalid id"
}

// UserNotFoundError — typed error для отсутствующего пользователя.
type UserNotFoundError struct {
	ID int
}

func (e UserNotFoundError) Error() string {
	// Текст ошибки не должен использоваться для принятия решений в программе.
	return "user not found"
}

// OpError добавляет контекст операции.
type OpError struct {
	Op  string
	Err error
}

func (e OpError) Error() string {
	// TODO: Сформируйте сообщение об ошибке в формате "Op: Err" (без потери исходной причины).
	// Важно: если Err == nil, верните разумный текст (например, Op).
	if e.Err == nil {
		return e.Op
	}
	return fmt.Sprintf("%s: %s", e.Op, e.Err)
}

func (e OpError) Unwrap() error {
	// TODO: Верните вложенную ошибку, чтобы errors.Unwrap мог доставать первопричину.
	return e.Err
}

func ParseID(raw string) (int, error) {
	id, err := strconv.Atoi(raw)
	if err != nil {
		// TODO: При ошибке парсинга возвращайте typed error InvalidIDError{Input: raw}.
		return 0, InvalidIDError{Input: raw}
	}
	return id, nil
}

func FindUser(users map[int]string, id int) (string, error) {
	name, ok := users[id]
	if !ok {
		// TODO: Если пользователя нет, возвращайте typed error UserNotFoundError{ID: id}.
		return "", UserNotFoundError{ID: id}
	}
	return name, nil
}

func GetUserName(users map[int]string, rawID string) (string, error) {
	// TODO: Вызовите ParseID и FindUser и оборачивайте любую ошибку в OpError, добавляя контекст операции.
	// Важно: цепочка errors.Unwrap должна доходить до InvalidIDError или UserNotFoundError.
	id, err := ParseID(rawID)
	if err != nil {
		return "", OpError{Op: "parse", Err: err}
	}

	name, err := FindUser(users, id)
	if err != nil {
		return "", OpError{Op: "findUser", Err: err}
	}

	return name, nil
}

func main() {
	users := map[int]string{
		1: "Ann",
		2: "Bob",
	}

	var rawID string
	fmt.Fscan(os.Stdin, &rawID)

	name, err := GetUserName(users, rawID)
	if err == nil {
		fmt.Println("NAME:", name)
		return
	}

	// TODO: Распознайте причину ошибки без парсинга текста:
	// пройдите по unwrap-цепочке (errors.Unwrap) и используйте type switch,
	// чтобы обработать InvalidIDError и UserNotFoundError.
	for err != nil {
		switch err.(type) {
		case InvalidIDError:
			fmt.Printf("INVALID ID: %s\n", err.(InvalidIDError).Input)
		case UserNotFoundError:
			fmt.Printf("NOT FOUND: %d\n", err.(UserNotFoundError).ID)
		default:
			fmt.Println("unknown error")
		}
		err = errors.Unwrap(err)
	}

}
