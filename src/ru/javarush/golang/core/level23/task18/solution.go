package main

/*
Вы пишете простейший блокнот: внутри у него список заметок, и для удобства вы встраиваете (embedded) этот список в
структуру блокнота, чтобы методы списка можно было вызывать прямо у блокнота. Но есть и обратная сторона: из-за embedding
наружу «выпирает» не только API методов, но и сам список как поле, и его можно случайно заменить целиком.

Напишите программу с типом Note (поле Text string) и типом NoteList (именованный тип на базе []Note). У NoteList должны
быть методы Add(text string) и Count() int. Add добавляет новую заметку в список.

Создайте тип Notebook с полем Title string и встроенным (embedded) NoteList. В main создайте блокнот, добавьте две заметки,
выведите строку before=2, затем напрямую замените весь список заметок через присваивание nb.NoteList = NoteList{} и
выведите строку after=0.

Программа должна выводить ровно:
before=2
after=0
*/

import "fmt"

type Note struct {
	Text string
}

// Именованный тип поверх слайса заметок.
type NoteList []Note

func (nl *NoteList) Add(text string) {
	// TODO: Реализуйте добавление новой заметки в список (через append).
	// TODO: Сохраните текст заметки в поле Note.Text.
	*nl = append(*nl, Note{text})
}

func (nl NoteList) Count() int {
	// TODO: Реализуйте подсчёт количества заметок в списке.
	return len(nl)
}

type Notebook struct {
	Title string
	// Embedding "протаскивает" наружу и методы, и само поле NoteList (его можно заменить целиком).
	NoteList
}

func main() {
	nb := Notebook{Title: "My notes"}

	// TODO: Добавьте в блокнот две заметки через nb.Add(...).
	nb.Add("first")
	nb.Add("second")

	fmt.Printf("before=%d\n", nb.Count())

	// TODO: Сбросьте список заметок присваиванием nb.NoteList = NoteList{}.
	nb.NoteList = NoteList{}

	fmt.Printf("after=%d\n", nb.Count())
}
