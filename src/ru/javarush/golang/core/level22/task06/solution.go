package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

// Pointer receiver нужен, чтобы менять исходный объект, а не копию.
func (p *Person) Birthday() {
	// TODO: Реализуйте изменение исходного объекта: увеличьте возраст (Age) на 1.
	p.Age++
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var personName string
	var personAge int
	fmt.Fscan(in, &personName, &personAge)

	p := Person{
		Name: personName,
		Age:  personAge,
	}

	p.Birthday() // ровно один вызов, без присваивания

	fmt.Println(p.Age)
}
