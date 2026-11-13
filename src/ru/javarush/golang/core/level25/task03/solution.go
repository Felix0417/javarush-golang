package main

/*
Вы делаете “пограничный лог” для консольного приложения: на входе вам дают название операции, её ID и флажок успеха.
Внутри логера вы хотите хранить всё это в одном контейнере map[string]any, чтобы позже легко расширять контекст новыми полями.

Программа должна прочитать из stdin три значения: operationName (строка без пробелов), requestID (целое число),
successFlag (0 или 1). Затем она должна построить map[string]any, где по ключу "op" лежит operationName, по ключу "id"
лежит requestID, а по ключу "success" лежит значение типа bool (это true, если successFlag == 1, иначе false).

После этого программа должна вывести две строки. В первой — значения в порядке op, затем id, затем success. Во второй —
типы этих значений (через %T) в том же порядке, с префиксом types:.
*/

import "fmt"

func main() {
	var operationName string
	var requestID int
	var successFlag int

	fmt.Scan(&operationName, &requestID, &successFlag)

	// Храним расширяемый контекст в одном контейнере.
	// TODO: сохраните название операции из входных данных по ключу "op"
	// TODO: сохраните ID запроса из входных данных по ключу "id"
	// TODO: преобразуйте флаг успеха из входных данных в bool и сохраните по ключу "success"
	m := map[string]any{
		"op":      operationName,
		"id":      requestID,
		"success": successFlag == 1,
	}

	// Значения нужно выводить в фиксированном порядке (не полагаясь на порядок map).
	fmt.Printf("%v %v %v\n", m["op"], m["id"], m["success"])
	fmt.Printf("types: %T %T %T\n", m["op"], m["id"], m["success"])
}
