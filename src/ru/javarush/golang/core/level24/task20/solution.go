package main

/*
Вы делаете тип ошибки для операций (например, "run", "open", "save"). И хотите аккуратно возвращать error из функций.
Но в Go есть известная ловушка: можно вернуть “typed nil” — когда внутри error лежит конкретный тип *OpError, равный nil,
и при этом сам интерфейс error оказывается не nil. В итоге err == nil внезапно становится false.

Определите тип ошибки OpError со строковым полем Op. Реализуйте интерфейс error методом Error() string, объявленным на
pointer receiver (*OpError).

Напишите две функции:
- BadDo(fail bool) error — при fail=false она возвращает переменную var e *OpError = nil как error (это создаёт typed nil).
- GoodDo(fail bool) error — при fail=false она возвращает именно nil (не typed nil).

При fail=true обе функции возвращают ненулевую ошибку (например, &OpError{Op: "run"}).

В main вызовите обе функции с fail=false и выведите результаты сравнений err == nil. Также выведите динамический тип (%T)
для каждой ошибки. Это ровно тот случай, о котором предупреждают в Go FAQ: возвращать конкретный тип ошибки вместо error —
частая причина неожиданных сравнений.
*/

import "fmt"

// OpError — типизированная ошибка операции.
//
// Важно: Error() должен быть объявлен на pointer receiver (*OpError),
// чтобы интерфейс error реализовывал именно *OpError (а не OpError).
type OpError struct {
	Op string
}

func (e *OpError) Error() string {
	// TODO: Реализуйте корректный текст ошибки (например, на основе поля Op).
	// Подумайте также, что делать, если e == nil.
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("error - %s", e.Op)
}

// BadDo демонстрирует ловушку "typed nil":
// при fail=false нужно вернуть *OpError, равный nil, как error (typed nil),
// чтобы err == nil неожиданно было false.
func BadDo(fail bool) error {
	if fail {
		return &OpError{Op: "run"}
	}

	// TODO: Верните typed nil: переменную вида `var e *OpError = nil` как `error`.
	var e *OpError = nil
	return e
}

// GoodDo возвращает "настоящий nil" как error при fail=false (без typed nil).
func GoodDo(fail bool) error {
	if fail {
		return &OpError{Op: "run"}
	}

	// TODO: Верните именно nil (не typed nil) при fail=false.
	if !fail {
		return nil
	}
	return &OpError{Op: ""}
}

func main() {
	bad := BadDo(false)
	good := GoodDo(false)

	fmt.Printf("BadDo(false): err == nil -> %v, type -> %T\n", bad == nil, bad)
	fmt.Printf("GoodDo(false): err == nil -> %v, type -> %T\n", good == nil, good)
}
