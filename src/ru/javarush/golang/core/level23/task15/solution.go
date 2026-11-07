package main

/*
Вы пишете примитивную систему тегов: один тип ставит метку "LEFT", другой — "RIGHT". Вы решили встроить оба в одну структуру,
чтобы получить их методы «как бы напрямую». Но тут всплывает важная особенность Go: если у двух embedded-типов есть метод
с одинаковым именем, то promoted-вызов становится неоднозначным, и придётся обращаться явно к каждому.

Создайте два типа LeftTag и RightTag. У каждого должен быть метод Tag() string, возвращающий разные строки: LeftTag.Tag()
возвращает "LEFT", а RightTag.Tag() возвращает "RIGHT".

Создайте тип Pair, который встраивает LeftTag и RightTag. В main создайте значение pair и выведите результат вызова обоих
методов одной строкой: "LEFT RIGHT". Из-за конфликта методов нельзя вызывать pair.Tag(), нужно использовать pair.LeftTag.Tag()
и pair.RightTag.Tag().
*/

import "fmt"

type LeftTag struct{}

func (LeftTag) Tag() string {
	// TODO: Реализуйте метод Tag() для LeftTag по условию задачи.
	return "LEFT"
}

type RightTag struct{}

func (RightTag) Tag() string {
	// TODO: Реализуйте метод Tag() для RightTag по условию задачи.
	return "RIGHT"
}

type Pair struct {
	LeftTag
	RightTag
}

func main() {
	pair := Pair{}

	// TODO: Исправьте вывод так, чтобы вызывались оба метода Tag() явно через embedded-типы.
	fmt.Printf("%s %s\n", pair.LeftTag.Tag(), pair.RightTag.Tag())
}
