package main

/*
Вы пишете “фабрику ресурсов” для приложения: в одном режиме создаём временный ресурс, в другом — говорим, что ресурса нет,
а в третьем — возвращаем ошибку. Самое коварное место: режим "нет ресурса" должен возвращать настоящий nil интерфейса,
чтобы проверка r == nil в main работала ожидаемо.

Объявите интерфейс Resource с методом Close() error. Реализуйте структуру TempResource с полем ID int и методом Close() error
(метод возвращает nil и может печатать closed <id>, либо может не печатать — но результат закрытия должен быть виден через main).

Напишите функцию NewResource(mode string) (Resource, error):
- mode == "temp": возвращается &TempResource{ID: 1}, error = nil;
- mode == "none": возвращается (nil, nil) так, чтобы в main выражение r == nil было true;
- любой другой mode: возвращается (nil, ошибка).

В main прочитайте mode из stdin, вызовите NewResource, выведите диагностику r через %T и %v, затем:
если r == nil, напечатайте NO RESOURCE;
иначе вызовите r.Close() и напечатайте CLOSE OK или CLOSE ERROR: ....
*/

import (
	"fmt"
)

type Resource interface {
	Close() error
}

type TempResource struct {
	ID int
}

func (r *TempResource) Close() error {
	return nil
}

func NewResource(mode string) (Resource, error) {
	switch mode {
	case "temp":
		// TODO: В режиме "temp" должен возвращаться реальный (не nil) ресурс с корректным ID.
		return &TempResource{ID: 1}, nil
	case "none":
		// TODO: В режиме "none" нужно вернуть именно nil-интерфейс (чтобы r == nil было true).
		return nil, nil
	default:
		// TODO: Для неизвестного режима нужно вернуть nil Resource и ненулевую ошибку.
		return nil, fmt.Errorf("unsupported mode: %s", mode)
	}
}

func main() {
	var mode string
	fmt.Scan(&mode)

	r, err := NewResource(mode)

	// Диагностика именно интерфейсного значения
	fmt.Printf("%T %v\n", r, r)

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	if r == nil {
		fmt.Println("NO RESOURCE")
		return
	}

	if err := r.Close(); err != nil {
		fmt.Printf("CLOSE ERROR: %v\n", err)
		return
	}

	fmt.Println("CLOSE OK")
}
