package main

import (
	"bufio"
	"fmt"
	"os"
)

type Counter struct {
	value int
}

// NewCounter — функция-конструктор (не метод), как требует условие.
func NewCounter(start int) Counter {
	// TODO: Инициализируйте счётчик значением start (а не значением по умолчанию).
	return Counter{value: start}
}

// Add и Inc меняют состояние, поэтому pointer receiver.
func (c *Counter) Add(delta int) {
	// TODO: Реализуйте добавление delta к текущему значению счётчика.
	c.value += delta
}

func (c *Counter) Inc() {
	// TODO: Реализуйте увеличение счётчика на 1.
	c.Add(1)
}

// Value только читает значение (без изменений).
func (c *Counter) Value() int {
	// TODO: Верните текущее значение счётчика, не изменяя его.
	return c.value
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var start, delta int
	fmt.Fscan(in, &start, &delta)

	counter := NewCounter(start)
	counter.Add(delta)
	counter.Inc()

	fmt.Println(counter.Value())
}
