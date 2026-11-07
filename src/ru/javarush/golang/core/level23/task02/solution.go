package main

import "fmt"

/*
Представьте, что вы пишете консольную заготовку для заметочника. Каждая заметка — это текст и «карточка» с ID, которую
	вы прячете внутрь структуры как embedded-часть, чтобы потом удобно обращаться к ID короткой записью.

Опишите структуру Meta с полем ID (int) и структуру Note с полем Text (string) и embedded-полем Meta. Программа читает
	из stdin два значения: целое число noteID и строку noteText (одно слово, без пробелов).

После чтения создайте переменную note типа Note через именованный литерал так, чтобы ID оказался внутри embedded-части,
	а Text — в основном поле Note. Затем выведите одну строку строго в формате Note <id>: <text>, причём ID возьмите коротким
	доступом note.ID (а не note.Meta.ID).
*/

// Meta — "карточка" с ID.
type Meta struct {
	ID int
}

// Note — заметка с текстом и метаданными.
type Note struct {
	Text string
	// TODO: Сделайте Meta embedded-полем (анонимным), чтобы ID был доступен как note.ID
	Meta
}

func main() {
	var noteID int
	var noteText string

	fmt.Scan(&noteID, &noteText)

	// TODO: Создайте переменную note типа Note через именованный литерал
	// TODO: Явно заполните Text и метаданные (ID должен попасть внутрь Meta)
	note := Note{Text: noteText, Meta: Meta{ID: noteID}}

	// TODO: Выведите ровно одну строку строго в формате: Note <id>: <text>
	// TODO: ID должен браться коротким доступом note.ID (а не note.Meta.ID)
	fmt.Printf("Note %d: %s", note.ID, note.Text)
}
