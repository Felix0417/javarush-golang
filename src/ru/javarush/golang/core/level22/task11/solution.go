package main

import "fmt"

// Enabler — минимальный контракт: "включить" и "проверить, включено ли".
type Enabler interface {
	Enable()
	IsOn() bool
}

type Switch struct {
	On bool
}

// Enable должен включать тумблер.
// TODO: Реализуйте включение тумблера так, чтобы изменение состояния сохранялось у исходного объекта (подумайте про receiver).
func (s *Switch) Enable() {
	// намеренно пусто
	if !s.On {
		s.On = true
	}
}

// IsOn возвращает текущее состояние тумблера.
func (s Switch) IsOn() bool {
	return s.On
}

func Activate(e Enabler) {
	// TODO: Включите устройство через интерфейс и выведите (Println) состояние "включено ли" через интерфейс.
	e.Enable()
	fmt.Println(e.IsOn())
}

func main() {
	var sw Switch

	// TODO: Передайте в Activate корректный аргумент, чтобы изменения внутри Enable() влияли на sw.
	Activate(&sw)

	fmt.Println(sw.IsOn())
}
