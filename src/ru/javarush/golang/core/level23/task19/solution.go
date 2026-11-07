package main

/*
Вы делаете «безопасный блокнот»: приложение должно отдавать заметки наружу, но так, чтобы внешний код не мог подменить
содержимое внутри блокнота. Иначе любой, кто получил слайс, мог бы тихо переписать вашу первую заметку на что-то
вроде "HACK". Значит, метод выдачи заметок обязан возвращать копию данных.

Напишите программу с типами Note (Text string) и NoteList (type NoteList []Note) с методами Add(text string) и Count() int.

Создайте тип SafeNotebook с полем Title string и неэкспортируемым полем notes типа NoteList (поле notes не embedded и
не экспортируемое). Реализуйте методы:
AddNote(text string) — добавляет заметку во внутренний список
Notes() []Note — возвращает копию данных (новый слайс, заполненный через copy)
Count() int — возвращает количество заметок

В main добавьте две заметки: Buy и Milk. Получите слайс ext := nb.Notes(), измените ext[0].Text на HACK и выведите:
internal_first=Buy
external_first=HACK
internal_count=2
*/

import "fmt"

type Note struct {
	Text string
}

type NoteList []Note

func (nl *NoteList) Add(text string) {
	*nl = append(*nl, Note{Text: text})
}

func (nl NoteList) Count() int {
	return len(nl)
}

type SafeNotebook struct {
	Title string
	notes NoteList
}

func (nb *SafeNotebook) AddNote(text string) {
	nb.notes.Add(text)
}

func (nb SafeNotebook) Notes() []Note {
	// TODO: Верните копию заметок (нужен новый слайс нужной длины и копирование элементов),
	// TODO: чтобы изменения во внешнем слайсе не влияли на внутреннее состояние блокнота.
	var copyNotes = make([]Note, len(nb.notes))
	copy(copyNotes, nb.notes)
	return copyNotes
}

func (nb SafeNotebook) Count() int {
	return nb.notes.Count()
}

func main() {
	nb := SafeNotebook{Title: "My notes"}

	nb.AddNote("Buy")
	nb.AddNote("Milk")

	ext := nb.Notes()
	ext[0].Text = "HACK"

	internalFirst := nb.Notes()[0].Text

	fmt.Println("internal_first=" + internalFirst)
	fmt.Println("external_first=" + ext[0].Text)
	fmt.Printf("internal_count=%d\n", nb.Count())
}
