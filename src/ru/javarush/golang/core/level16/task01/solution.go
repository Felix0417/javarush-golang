package main

import "fmt"

func main() {
	var err error

	// TODO: Выведите результат сравнения err == nil (только этим выражением).
	fmt.Printf("err_is_nil=%v\n", err == nil)

	// TODO: Выведите значение переменной err.
	fmt.Printf("err_value=%v\n", err)

	// TODO: Выведите тип переменной err через fmt.Printf и плейсхолдер %T.
	fmt.Printf("err_type=%T\n", err)
}
