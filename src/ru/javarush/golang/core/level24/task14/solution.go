package main

/*
Представьте, что вы делаете “сервис заметок”: бизнес‑логика хочет уметь сохранять заметку и получать весь список. Но
сервис не должен зависеть от того, где именно хранятся данные — в памяти, в файле или в базе. Поэтому вы задаёте
минимальный внутренний интерфейс прямо возле сервиса (и не экспортируете его наружу).

Опишите структуру NoteService, которая хранит поле st типа storage (интерфейс). Интерфейс storage должен быть
неэкспортируемым (с маленькой буквы) и объявлен рядом с NoteService. В интерфейсе должны быть методы Save(text string) error
и All() []string — ничего лишнего.

Сделайте конкретную реализацию memStorage, которая хранит заметки в слайсе строк. Методы Save и All должны быть
реализованы так, чтобы сервис мог добавить две заметки и вывести их построчно в stdout в том же порядке, в котором
они добавлялись.
*/

import "fmt"

// storage — внутренний контракт хранилища (не экспортируем наружу).
// Важно: ровно два метода, как требует условие.
type storage interface {
	Save(text string) error
	All() []string
}

// NoteService зависит только от интерфейса storage.
type NoteService struct {
	st storage
}

// NewNoteService — конструктор с внедрением зависимости.
func NewNoteService(st storage) *NoteService {
	return &NoteService{st: st}
}

// Save — бизнес-метод сервиса: сохраняем заметку.
func (s *NoteService) Save(text string) error {
	// TODO: Сохраните заметку через хранилище s.st (без прямой зависимости от конкретного типа).
	return s.st.Save(text)
}

// All — бизнес-метод сервиса: получаем все заметки.
func (s *NoteService) All() []string {
	// TODO: Верните список заметок через хранилище s.st.
	return s.st.All()
}

// memStorage — реализация storage в памяти на базе []string.
type memStorage struct {
	notes []string
}

func (m *memStorage) Save(text string) error {
	// TODO: Добавьте заметку text в m.notes так, чтобы сохранялся порядок добавления.
	if m.notes == nil {
		m.notes = make([]string, 0)
	}
	m.notes = append(m.notes, text)
	return nil
}

func (m *memStorage) All() []string {
	// TODO: Верните все заметки в порядке добавления.
	return m.notes
}

func main() {
	st := &memStorage{}
	svc := NewNoteService(st)

	_ = svc.Save("Первая заметка")
	_ = svc.Save("Вторая заметка")

	for _, note := range svc.All() {
		fmt.Println(note)
	}
}
