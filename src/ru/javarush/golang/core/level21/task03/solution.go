package main

import "fmt"

// Counter — счётчик очков игрока.
type Counter struct {
	Name  string
	Value int
}

// Add увеличивает Value на delta прямо в переданной структуре.
func Add(c *Counter, delta int) {
	// TODO: увеличьте значение счётчика (поле Value) на delta, изменяя структуру по указателю
	c.Value += delta
}

func main() {
	var playerName string
	var firstDelta, secondDelta int
	fmt.Scan(&playerName, &firstDelta, &secondDelta)

	// По требованию: создаём через var, потом записываем имя.
	var c Counter
	// TODO: запишите имя игрока в c.Name
	c.Name = playerName

	// TODO: вызовите Add ровно два раза и строго в порядке: сначала с firstDelta, затем с secondDelta
	Add(&c, firstDelta)
	Add(&c, secondDelta)

	// TODO: выведите одну строку формата Name=%q Value=%d (имя через %q)
	fmt.Printf("Name=%q Value=%d\n", c.Name, c.Value)
}
