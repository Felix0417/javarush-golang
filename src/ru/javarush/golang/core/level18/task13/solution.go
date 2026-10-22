package main

import "fmt"

// Демонстрация: recover() полезен только внутри deferred-функции во время раскрутки паники.
func boom() {
	// TODO: Добавьте defer func(){ ... }, внутри которого вызывается recover() и обрабатывается паника.
	// TODO: Внутри deferred-функции выведите в stdout строку строго формата: recovered: boom
	// TODO: После установки defer вызовите panic("boom"), чтобы запустить раскрутку паники.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered: boom")
		}
	}()

	panic("boom")
}

func main() {
	// Прямой recover без паники всегда возвращает nil.
	fmt.Printf("direct recover: %v\n", recover())

	boom()

	fmt.Println("after boom")
}
