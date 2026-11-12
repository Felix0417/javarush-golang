package main

/*
Вы делаете простейшие заметки “в оперативке”: можно добавить текст, можно показать список. Но вы хотите, чтобы разные
части программы просили только то, что им нужно: функция добавления не должна требовать умения “листать”, а функция
показа списка не должна требовать умения “добавлять”. Так получается гибче и проще поддерживать.

Создайте структуру Note с полями ID int и Text string. Опишите интерфейс NoteAdder с методом Add(text string) (Note, error) и
интерфейс NoteLister с методом List() []Note. Реализуйте тип MemoryNotes со слайсом заметок и счётчиком nextID.
Метод Add должен возвращать ошибку, если text — пустая строка (текст ошибки строго "text is empty"); иначе добавлять
заметку с новым ID. Метод List возвращает текущий список в порядке добавления.

Напишите функции RunAdd(a NoteAdder, text string) и RunList(l NoteLister).
В main выполните:
RunAdd(store, "buy milk")
RunAdd(store, "")
RunAdd(store, "learn interfaces")
RunList(store)

Ожидаемый вывод:
added: 1 buy milk
add error: text is empty
added: 2 learn interfaces
1 buy milk
2 learn interfaces
*/

import (
	"errors"
	"fmt"
)

type Note struct {
	ID   int
	Text string
}

// NoteAdder — маленький интерфейс только для добавления.
type NoteAdder interface {
	Add(text string) (Note, error)
}

func (n Note) String() string {
	return fmt.Sprintf("%d %s\n", n.ID, n.Text)
}

// NoteLister — маленький интерфейс только для просмотра списка.
type NoteLister interface {
	List() []Note
}

type MemoryNotes struct {
	notes  []Note
	nextID int
}

func (m *MemoryNotes) Add(text string) (Note, error) {
	// TODO: Реализуйте валидацию: если text пустой — вернуть ошибку с требуемым текстом.
	if text == "" {
		return Note{}, errors.New("text is empty")
	}

	// TODO: Реализуйте генерацию ID (последовательно) и сохранение заметки в памяти (append в слайс).
	m.nextID++
	n := Note{
		ID:   m.nextID,
		Text: text,
	}
	m.notes = append(m.notes, n)

	return n, nil
}

func (m *MemoryNotes) List() []Note {
	// TODO: Верните текущий список заметок в порядке добавления.
	return m.notes
}

func RunAdd(a NoteAdder, text string) {
	// TODO: Реализуйте вывод результата добавления в требуемом формате.
	id, err := a.Add(text)
	if err != nil {
		fmt.Printf("add error: %s\n", err.Error())
		return
	}
	fmt.Printf("added: %s", id)
}

func RunList(l NoteLister) {
	// TODO: Реализуйте вывод списка заметок (каждая заметка с новой строки) в требуемом формате.
	for _, n := range l.List() {
		fmt.Print(n.String())
	}
}

func main() {
	store := &MemoryNotes{}
	// TODO: Инициализируйте nextID так, чтобы первая добавленная заметка получила корректный ID.
	store.nextID = 0

	RunAdd(store, "buy milk")
	RunAdd(store, "")
	RunAdd(store, "learn interfaces")
	RunList(store)
}
