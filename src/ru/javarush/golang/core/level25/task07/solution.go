package main

/*
Вы пишете кусочек кода для профилей: иногда профиль приходит как значение User, иногда как указатель *User (например, из кеша),
а иногда приходит “указатель нужного типа, но пустой” (nil). Вам важно не перепутать эти случаи: если внутри лежит *User == nil,
имя извлекать нельзя, даже если ok у type assertion равен true.

Программа должна объявить структуру User с полем Name string. Из stdin читаются два токена: mode и name. mode может
быть "value", "ptr" или "nilptr".

В зависимости от mode создайте переменную payload any: при "value" положите User{Name: name}, при "ptr" положите
&User{Name: name}, при "nilptr" положите значение типа *User, равное nil.

Реализуйте extractName(x any) (string, bool): сначала попробуйте безопасно извлечь имя из User, затем из *User. Если
внутри *User лежит nil, верните ("", false), даже если ok у assertion равен true. В main напечатайте либо name=<имя>, либо no name.

Программа должна объявить структуру User с полем Name string.
Из stdin читаются два токена: mode и name.
mode может быть "value", "ptr" или "nilptr".

В зависимости от mode создайте переменную payload any:
при "value" положите User{Name: name},
при "ptr" положите &User{Name: name},
при "nilptr" положите значение типа *User, равное nil.

Реализуйте extractName(x any) (string, bool):
сначала попробуйте безопасно извлечь имя из User, затем из *User.
Если внутри *User лежит nil, верните ("", false),
даже если ok у assertion равен true.
В main напечатайте либо name=<имя>, либо no name.
*/

import (
	"fmt"
	"os"
)

type User struct {
	Name string
}

func extractName(x any) (string, bool) {
	// 1) Сначала пробуем извлечь значение User.
	if u, ok := x.(User); ok {
		return u.Name, true
	}

	// 2) Затем пробуем извлечь *User.
	// TODO: Реализуйте извлечение имени из *User через безопасный type assertion (v, ok := x.(*User)).
	// TODO: Учтите случай, когда внутри интерфейса лежит (*User)(nil): в этом случае нельзя читать поле Name.
	if v, ok := x.(*User); ok && v != nil {
		return v.Name, true
	}

	return "", false
}

func main() {
	var mode, name string
	fmt.Fscan(os.Stdin, &mode, &name)

	var payload any
	switch mode {
	case "value":
		payload = User{Name: name}
	case "ptr":
		payload = &User{Name: name}
	case "nilptr":
		// TODO: Положите в payload значение типа *User, равное nil (типизированный nil), а не nil-интерфейс.
		var usr *User = nil
		payload = usr
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
		return
	}

	if n, ok := extractName(payload); ok {
		fmt.Printf("name=%s\n", n)
	} else {
		fmt.Println("no name")
	}
}
