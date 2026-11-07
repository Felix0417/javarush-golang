package main

import (
	"fmt"
	"os"
)

/*
Вы пишете in-memory хранилище заметок для маленького консольного приложения. Команда хочет, чтобы хранилище умело и
добавлять заметки, и искать их по ID. Вместо одного «жирного» интерфейса вы описываете два маленьких и собираете общий
контракт через embedding, чтобы было видно: что именно нужно на каждом шаге.

Объявите структуру Note с полями ID (int) и Text (string). Объявите интерфейс NoteAdder с методом Add(n Note) error и
интерфейс NoteFinder с методом FindByID(id int) (Note, bool). Объявите интерфейс NoteStore, который включает оба этих
интерфейса через embedding.

Реализуйте тип MemNoteStore (например, со слайсом заметок), который удовлетворяет NoteStore. Add должен возвращать ошибку,
если заметка с таким ID уже есть. Программа читает из stdin 3 заметки (пары ID и Text без пробелов), затем читает queryID
и печатает либо found: <id> <text>, либо not found. Если при добавлении была ошибка, программа печатает
строку error: <текст ошибки>.
*/

type Note struct {
	ID   int
	Text string
}

type NoteAdder interface {
	Add(n Note) error
}

type NoteFinder interface {
	FindByID(id int) (Note, bool)
}

type NoteStore interface {
	NoteAdder
	NoteFinder
}

type MemNoteStore struct {
	notes []Note
}

func (s *MemNoteStore) Add(n Note) error {
	// TODO: Реализуйте добавление заметки в in-memory хранилище.
	// TODO: Проверьте уникальность ID: если такой ID уже есть в хранилище — верните error (panic использовать нельзя).
	for _, v := range s.notes {
		if v.ID == n.ID {
			return fmt.Errorf("note with id-%d already exists", n.ID)
		}
	}

	s.notes = append(s.notes, n)
	return nil
}

func (s *MemNoteStore) FindByID(id int) (Note, bool) {
	// TODO: Реализуйте поиск заметки по заданному id по всему хранилищу.
	// TODO: Верните (Note, true) при успешном нахождении, иначе (нулевую Note, false).

	if len(s.notes) == 0 {
		return Note{}, false
	}

	for i, n := range s.notes {
		if n.ID == id {
			return s.notes[i], true
		}
	}
	return Note{}, false
}

func main() {
	in := os.Stdin

	var store NoteStore = &MemNoteStore{}

	for i := 0; i < 3; i++ {
		var id int
		var text string
		fmt.Fscan(in, &id, &text)

		if err := store.Add(Note{ID: id, Text: text}); err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}

	var queryID int
	fmt.Fscan(in, &queryID)

	if n, ok := store.FindByID(queryID); ok {
		fmt.Printf("found: %d %s\n", n.ID, n.Text)
	} else {
		fmt.Println("not found")
	}
}
