package main

/*
Вы делаете консольный вывод для тикетов поддержки. У тикета есть заголовок, а также два набора идентификаторов: один
«основной», другой — из аудита, плюс имя того, кто создал запись. Вы встраиваете обе структуры, но так как у них
пересекаются имена полей (ID), в методе придётся быть аккуратным и всегда писать явный путь.

Создайте типы Ident и Audit, у каждого поле ID типа int. У Audit также есть поле CreatedBy типа string.
Создайте тип Ticket, который содержит поле Title типа string и встраивает Ident и Audit. Реализуйте метод Summary() string
у Ticket. Он должен возвращать строку формата title=<Title> ident=<Ident.ID> audit=<Audit.ID> by=<Audit.CreatedBy>.

Программа читает из stdin: title (одно слово), identID (int), auditID (int), createdBy (одно слово). Затем создаёт Ticket,
заполняет все поля и печатает результат ticket.Summary().
*/

import "fmt"

type Ident struct {
	ID int
}

type Audit struct {
	ID        int
	CreatedBy string
}

type Ticket struct {
	Title string
	Ident // embedded
	Audit // embedded
}

func (t Ticket) Summary() string {
	// TODO: Реализуйте формирование строки сводки по заданному формату.
	// TODO: Из-за конфликта имён поля ID в embedded-структурах обращайтесь к полям только явными путями: t.Ident.ID и t.Audit.ID.
	// TODO: Строка должна строго соответствовать шаблону: title=<Title> ident=<Ident.ID> audit=<Audit.ID> by=<Audit.CreatedBy>
	return fmt.Sprintf("title=%s ident=%d audit=%d by=%s", t.Title, t.Ident.ID, t.Audit.ID, t.Audit.CreatedBy)
}

func main() {
	var title string
	var identID int
	var auditID int
	var createdBy string

	fmt.Scan(&title, &identID, &auditID, &createdBy)

	// TODO: Корректно заполните все поля ticket считанными значениями (Title, Ident.ID, Audit.ID, Audit.CreatedBy), используя явные пути.
	var ticket = Ticket{
		Title: title,
		Ident: Ident{ID: identID},
		Audit: Audit{ID: auditID, CreatedBy: createdBy},
	}

	fmt.Print(ticket.Summary())
}
