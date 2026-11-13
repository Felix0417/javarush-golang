package main

/*
Представьте консольного бота: ему скармливают команды разных типов, а он должен выдать понятный результат. Если команда
неизвестного типа — не паниковать, а вернуть нормальную ошибку с текстом, в котором видно, какой тип прилетел.

Создайте три типа команд: AddCmd с полем Title string, DoneCmd с полем ID int, ListCmd без полей. Напишите функцию
Run(cmd any) (string, error), которая через type switch формирует строку результата: для AddCmd — ADD:<title>, для
DoneCmd — DONE:<id>, для ListCmd — LIST. Для неизвестного типа нужно вернуть error с текстом unknown command type <type>,
где <type> — это %T от cmd.

В main создайте слайс []any из трёх элементов: AddCmd{Title:"Buy milk"}, DoneCmd{ID:3}, "oops". Для каждого элемента
вызовите Run. Если ошибка не nil — выведите ERR: + текст ошибки, иначе выведите строку результата.

Ожидаемый вывод:
ADD:Buy milk
DONE:3
ERR:unknown command type string
*/

import "fmt"

type AddCmd struct {
	Title string
}

type DoneCmd struct {
	ID int
}

type ListCmd struct{}

// Run обрабатывает команду по её типу через type switch.
func Run(cmd any) (string, error) {
	switch c := cmd.(type) {
	case AddCmd:
		// TODO: Сформируйте строку результата для AddCmd в требуемом формате.
		return fmt.Sprintf("ADD:%s", c.Title), nil
	case DoneCmd:
		// TODO: Сформируйте строку результата для DoneCmd в требуемом формате.
		return fmt.Sprintf("DONE:%d", c.ID), nil
	case ListCmd:
		// TODO: Сформируйте строку результата для ListCmd в требуемом формате.
		return fmt.Sprintf("LIST"), nil
	default:
		// Для неизвестного типа — обычная ошибка, без panic.
		return "", fmt.Errorf("unknown command type %T", cmd)
	}
}

func main() {
	cmds := []any{
		AddCmd{Title: "Buy milk"},
		DoneCmd{ID: 3},
		"oops",
	}

	for _, cmd := range cmds {
		out, err := Run(cmd)
		if err != nil {
			fmt.Println("ERR:" + err.Error())
			continue
		}
		fmt.Println(out)
	}
}
