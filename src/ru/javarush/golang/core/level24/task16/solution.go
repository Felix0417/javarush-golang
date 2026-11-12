package main

/*
У вас есть один “универсальный” репозиторий в памяти, но разные части программы используют его по‑разному. Один сценарий
хочет только добавлять и читать список, другой — очищать и читать список. В Go принято объявлять маленькие интерфейсы у
потребителей, чтобы контракт был минимальным и не разрастался “на всякий случай”.

Создайте тип MemRepo, который хранит элементы в памяти и имеет методы Add(text string) error, List() []string и Clear().

Сделайте две функции:
RunAddScenario(r …, a, b string), которая добавляет две строки и печатает весь список построчно.
RunClearScenario(r …), которая очищает репозиторий и печатает строку count=0.

Для каждой функции объявите отдельный интерфейс рядом с ней: первый должен требовать Add и List, второй — Clear и List.
Не делайте один большой интерфейс “комбайн”.
*/

import "fmt"

// MemRepo хранит строки в памяти.
type MemRepo struct {
	items []string
}

func (r *MemRepo) Add(text string) error {
	// TODO: Реализуйте добавление строки в репозиторий в памяти.
	if r.items == nil {
		r.items = []string{}
	}
	r.items = append(r.items, text)
	return nil
}

func (r *MemRepo) List() []string {
	// TODO: Верните текущий список строк в порядке добавления.
	return r.items
}

func (r *MemRepo) Clear() {
	// TODO: Очистите репозиторий (после этого List должна возвращать пустой список).
	clear(r.items)
	r.items = nil
}

// addRepo — минимальный контракт ровно для сценария добавления.
type addRepo interface {
	Add(string) error
	List() []string
}

func RunAddScenario(r addRepo, a, b string) {
	// TODO: Добавьте две строки (a и b) в репозиторий и выведите весь список построчно.
	_ = r.Add(a)
	_ = r.Add(b)

	for _, s := range r.List() {
		fmt.Println(s)
	}
}

// clearRepo — минимальный контракт ровно для сценария очистки.
type clearRepo interface {
	Clear()
	List() []string
}

func RunClearScenario(r clearRepo) {
	// TODO: Очистите репозиторий и выведите в stdout ровно строку count=0.
	r.Clear()
	fmt.Println("count=0")
}

func main() {
	// Пустой main: сценарии будут вызываться проверяющей системой.
}
