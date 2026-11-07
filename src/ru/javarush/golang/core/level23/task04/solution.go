package main

import "fmt"

/*
Вы поддерживаете простейшую систему документов. У каждого документа есть имя, а внутри — метаданные с ID и Version.
Когда документ проходит «перепроверку», вы должны увеличить версию на 1. Важно показать, что изменение версии через
короткий promoted-доступ (doc.Version) на самом деле меняет то же поле, что и доступ через doc.Meta.Version.

Опишите структуру Meta с полями ID (int) и Version (int). Опишите структуру Document с полем Name (string) и embedded-полем
Meta (значение, не указатель).

Напишите функцию BumpVersion, которая принимает указатель на Document и увеличивает версию на 1, используя короткий доступ
к promoted-полю: d.Version++.

В main программа читает из stdin три значения: id (int), version (int), name (одно слово). Затем создаёт Document через
именованный литерал, вызывает BumpVersion и выводит одну строку в формате version=<v1> metaVersion=<v2>, где v1 берётся
через doc.Version, а v2 — через doc.Meta.Version. Оба числа должны совпадать и быть на 1 больше исходного version.
*/

type Meta struct {
	ID      int
	Version int
}

type Document struct {
	Name string
	Meta // embedded как значение
}

func BumpVersion(d *Document) {
	// TODO: увеличьте версию документа на 1 через promoted-поле (короткий доступ)
	d.Version++
}

func main() {
	var id, version int
	var name string
	fmt.Scan(&id, &version, &name)

	doc := Document{
		// TODO: создайте документ через именованный литерал и заполните Meta данными из ввода
		Name: name,
		Meta: Meta{ID: id, Version: version},
	}

	BumpVersion(&doc)

	fmt.Printf("version=%d metaVersion=%d\n", doc.Version, doc.Meta.Version)
}
