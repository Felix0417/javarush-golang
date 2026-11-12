package main

/*
Представьте, что вы делаете “мешок чисел” для статистики: в него можно добавлять значения, а потом спрашивать, сколько
всего накопилось. Кажется простым — но здесь важно почувствовать, почему интерфейсы иногда требуют именно методы с
pointer receiver: добавление меняет состояние.

Опишите интерфейс IntBag с двумя методами: Add(x int) и Len() int. Реализуйте его типом MemBag, внутри которого хранится
слайс items []int.

Метод Add должен быть объявлен с pointer receiver (receiver = *MemBag) и добавлять элемент в слайс. Метод Len должен
возвращать текущую длину items и может быть объявлен с value receiver.

В main создайте переменную интерфейсного типа IntBag, которая ссылается на пустой MemBag через указатель. Добавьте
числа 10 и 20 и выведите строку вида: len=2 type=*main.MemBag. Ситуация “method … has pointer receiver” — типичная
подсказка компилятора, если попытаться положить в интерфейс не то.
*/

import "fmt"

// IntBag — контракт "мешка": можно добавлять и узнавать размер.
type IntBag interface {
	Add(x int)
	Len() int
}

// MemBag хранит значения в памяти.
type MemBag struct {
	items []int
}

// Add обязан быть с pointer receiver: метод меняет состояние (слайс в структуре).
func (b *MemBag) Add(x int) {
	// TODO: Реализуйте добавление элемента x в b.items (с изменением состояния).
	if b.items == nil {
		b.items = make([]int, 0)
	}
	b.items = append(b.items, x)
}

// Len состояние не меняет, поэтому можно value receiver.
func (b MemBag) Len() int {
	// TODO: Верните текущую длину b.items.
	return len(b.items)
}

func main() {
	// В интерфейс кладём именно *MemBag, иначе не будет метода Add (pointer receiver).
	var bag IntBag = &MemBag{}

	bag.Add(10)
	bag.Add(20)

	fmt.Printf("len=%d type=%T", bag.Len(), bag)
}
