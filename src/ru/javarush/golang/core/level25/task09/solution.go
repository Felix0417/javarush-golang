package main

/*
Вы делаете мини-инструмент для отладочной панели: ему передают any, а он должен коротко подписать, что внутри лежит. Для
нескольких типов подписи фиксированные, а если прилетело что-то неожиданное — показать unknown и реальный тип.

Напишите функцию Describe(x any) string, которая возвращает строку-описание значения внутри x, используя type switch.
Функция должна обрабатывать int, string, bool. Для неизвестных типов функция должна возвращать строку формата unknown:<type>,
где <type> — это результат %T для исходного x.

В main создайте четыре значения типа any: 7, "go", true, 3.14. Вызовите Describe для каждого и выведите результат построчно.

Ожидаемый вывод:
int:7
string:go
bool:true
unknown:float64
*/

import "fmt"

func Describe(x any) string {
	// Требование: type switch строго в виде switch v := x.(type)
	switch v := x.(type) {
	case int:
		// TODO: Верните строку-описание для значения типа int в требуемом формате, используя v.
		return fmt.Sprintf("%T:%d", v, v)
	case string:
		// TODO: Верните строку-описание для значения типа string в требуемом формате, используя v.
		return fmt.Sprintf("%T:%s", v, v)
	case bool:
		// TODO: Верните строку-описание для значения типа bool в требуемом формате, используя v.
		return fmt.Sprintf("%T:%t", v, v)
	default:
		// TODO: Верните строку для неизвестного типа; тип определяйте через %T от исходного x (не от v).
		return fmt.Sprintf("unknown:%T", x)
	}
}

func main() {
	var a any = 7
	var b any = "go"
	var c any = true
	var d any = 3.14

	fmt.Println(Describe(a))
	fmt.Println(Describe(b))
	fmt.Println(Describe(c))
	fmt.Println(Describe(d))
}
