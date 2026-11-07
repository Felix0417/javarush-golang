package main

/*
Вы делаете структуру документа, у которого есть своя версия, но внутри ещё лежит встроенная мета-информация (например,
версия формата). И в какой-то момент команда решила назвать поле одинаково: Version и снаружи, и внутри Meta. В Go внешнее
поле «побеждает» promoted-поле, поэтому короткий доступ d.Version будет читать версию документа, а до версии в Meta придётся
идти явно.

Описан тип Meta с полем Version типа int. Создайте тип Document, который встраивает Meta (embedding) и дополнительно
содержит собственное поле Version типа int (то же имя, что и в Meta).
Программа читает из stdin два целых числа: docVersion и metaVersion. Затем создаёт Document так, чтобы Document.Version
был равен docVersion, а Document.Meta.Version был равен metaVersion.

Программа выводит одну строку: doc=<docVersion> meta=<metaVersion> diff=<docVersion-metaVersion>. Для версии внутри меты
используйте явный доступ d.Meta.Version.
*/

import "fmt"

type Meta struct {
	Version int
}

type Document struct {
	Meta
	// TODO: Добавьте собственное поле Version типа int, чтобы оно перекрывало Meta.Version при обращении d.Version.
	Version int
}

func main() {
	var docVersion, metaVersion int
	fmt.Scan(&docVersion, &metaVersion)

	d := Document{}

	// TODO: Инициализируйте Document так, чтобы d.Version было равно docVersion, а d.Meta.Version было равно metaVersion.
	d.Version = docVersion
	d.Meta.Version = metaVersion

	// TODO: Выведите строку в нужном формате и посчитайте diff как разницу версии документа и версии в мета-информации.
	// TODO: Для версии внутри мета-информации используйте явный доступ d.Meta.Version.
	fmt.Printf("doc=%d meta=%d diff=%d", d.Version, d.Meta.Version, d.Version-d.Meta.Version)
}
