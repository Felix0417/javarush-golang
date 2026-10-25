package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	ID       int    `json:"id" db:"order_id"`            // TODO: задайте значения json и db тегов для поля ID
	Customer string `json:"customer" db:"customer_name"` // TODO: задайте значения json и db тегов для поля Customer
	Paid     bool   `json:"paid" db:"is_paid"`           // TODO: задайте значения json и db тегов для поля Paid
}

func main() {
	t := reflect.TypeOf(Order{})

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		// TODO: прочитайте значение json-тега строго через f.Tag.Get("json")
		jsonTag := f.Tag.Get("json")

		// TODO: прочитайте значение db-тега строго через f.Tag.Get("db")
		dbTag := f.Tag.Get("db")

		// Формат вывода должен быть строго таким, как требуется в задании.
		fmt.Printf("%s json=%s db=%s\n", f.Name, jsonTag, dbTag)
	}
}
