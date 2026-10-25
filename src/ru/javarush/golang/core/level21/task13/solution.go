package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	// TODO: Добавьте struct tag с ключом json для поля Title.
	Title string `json:"title"`
	// TODO: Добавьте struct tag с ключом json для поля Pages.
	Pages int `json:"pages"`
}

func main() {
	// Тип нужно получать именно через reflect.TypeOf(Book{}).
	t := reflect.TypeOf(Book{})

	// TODO: Обработайте ситуацию, когда поле не найдено (ok == false).
	titleField, ok := t.FieldByName("Title")
	if !ok {
		fmt.Errorf("title field not found")
	}
	// TODO: Обработайте ситуацию, когда поле не найдено (ok == false).
	pagesField, ok := t.FieldByName("Pages")
	if !ok {
		fmt.Errorf("pages field not found")
	}

	// Печатаем значения json-тегов (пока они пустые, пока студент не добавит теги).
	fmt.Println(titleField.Tag.Get("json"))
	fmt.Println(pagesField.Tag.Get("json"))
}
