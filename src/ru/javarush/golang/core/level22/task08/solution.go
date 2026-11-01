package main

import "fmt"

// Note — модель заметки (поля строго по требованию).
type Note struct {
	Title string
	Done  bool
	Tags  []string
}

// NewNote возвращает указатель, чтобы дальше строить fluent-цепочку.
func NewNote(title string) *Note {
	// TODO: Инициализируйте заметку: установите Title из аргумента и верните *Note.
	return &Note{Title: title}
}

// AddTag — fluent-метод с pointer receiver.
func (n *Note) AddTag(tag string) *Note {
	// TODO: Добавьте tag в слайс Tags (через append) и верните тот же *Note для продолжения цепочки.
	n.Tags = append(n.Tags, tag)
	return n
}

// MarkDone — fluent-метод с pointer receiver.
func (n *Note) MarkDone() *Note {
	// TODO: Установите Done = true и верните тот же *Note для продолжения/завершения цепочки.
	n.Done = true
	return n
}

func main() {
	var noteTitle, tagOne, tagTwo string
	fmt.Scan(&noteTitle, &tagOne, &tagTwo)

	// Цепочка без промежуточной переменной до завершения цепочки.
	note := NewNote(noteTitle).
		AddTag(tagOne).
		AddTag(tagTwo).
		MarkDone()

	fmt.Println(note.Done)
	fmt.Println(len(note.Tags))
}
