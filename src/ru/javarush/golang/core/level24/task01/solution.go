package main

import "fmt"

// Greeter — маленький контракт: кто умеет здороваться.
type Greeter interface {
	Greet() string
}

// User — конкретная реализация приветствия.
type User struct {
	Name string
}

// Greet реализует интерфейс Greeter неявно (по совпадению сигнатуры).
func (u User) Greet() string {
	// TODO: Сформируйте приветствие, используя имя пользователя u.Name, и верните готовую строку без лишних пробелов в конце.
	return fmt.Sprintf("Hello, %s!", u.Name)
}

func main() {
	var name string
	fmt.Scan(&name)

	var g Greeter
	g = User{Name: name} // используем User через интерфейс

	fmt.Println(g.Greet())
}
