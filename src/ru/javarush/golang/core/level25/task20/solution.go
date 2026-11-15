package main

/*
Вы собираете маленький “шлюз команд” для консольного трекера задач. Команды приходят разные, ошибки тоже разные, но вывод
должен быть стабильным: для пользователя — понятный код и сообщение, для разработчика — полный DEBUG с технической
цепочкой. И самое важное: тип ошибки нужно определять корректно через errors.As, а не через поиск подстрок в тексте.

Напишите программу, которая читает из stdin команду.

Поддерживаются команды: add <title> (где title — одно слово, а title == "-" считается некорректным) и get <id>.

В программе заранее есть один “хранимый” таск: ID=1, Title=demo, Done=false.

Определите typed errors: ValidationError (value) с полями Field, Msg; NotFoundError (pointer) с полем ID и методом
Error() на *NotFoundError; OpError (pointer) с Op, Err и Unwrap.

Реализуйте функции AddTask и GetTask, которые при ошибках возвращают &OpError{Op: "...", Err: <причина>}.

Реализуйте ToUserError(err error) (code, message string, field, detail string): если внутри есть validation, вернуть
code=validation, message=Некорректные данные, а field/detail взять из ValidationError; если внутри есть not_found,
вернуть code=not_found, message=Задача не найдена; иначе вернуть code=internal, message=Не удалось выполнить операцию.

При ошибке программа должна выводить:
CODE=<code>
MESSAGE=<message>
FIELD=<field> DETAIL=<detail>   (только для validation; для остальных не печатать)
DEBUG=<err.Error()>

При успехе:
для add вывести OK id=<id>
для get вывести TASK id=<id> title=<title> done=<done>
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//  — цепочки ошибок + errors.As по всей цепочке

type Task struct {
	ID    int
	Title string
	Done  bool
}

// ValidationError — value-тип (не указатель)
type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	// TODO: Сформируйте стабильный текст ошибки валидации, используя Field и Msg.
	// TODO: Учитывайте, что DEBUG печатает err.Error(), поэтому строка должна быть информативной.
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// NotFoundError — используется как pointer-тип
type NotFoundError struct {
	ID int
}

func (e *NotFoundError) Error() string {
	// TODO: Сформируйте стабильный текст ошибки "не найдено", включая ID задачи.
	// TODO: Метод должен быть на *NotFoundError (pointer receiver).
	if e == nil {
		return "task not found"
	}
	return fmt.Sprintf("не найдено %d", e.ID)
}

// OpError — обёртка операции, чтобы DEBUG показывал техническую цепочку
type OpError struct {
	Op  string
	Err error
}

func (e *OpError) Error() string {
	// TODO: Сформируйте текст ошибки так, чтобы в DEBUG было видно операцию и первопричину.
	// TODO: Учтите, что ошибка может быть обёрнута в несколько слоёв.
	if e == nil {
		return "operation error"
	}
	return fmt.Sprintf("%s: %v", e.Op, e.Err)
}

func (e *OpError) Unwrap() error {
	// TODO: Верните первопричину (Err), чтобы errors.As проходил по цепочке ошибок.
	return e.Err
}

var demoTask = Task{ID: 1, Title: "demo", Done: false}

func AddTask(title string) (int, error) {
	// TODO: Реализуйте команду add.
	// TODO: Если title == "-", верните ошибку, обёрнутую в *OpError (Op="AddTask"),
	// TODO: а причиной должна быть typed error ValidationError (value-тип) с Field="title" и деталью в Msg.
	if title == "-" {
		return 0, &OpError{Op: "AddTask", Err: ValidationError{Field: "title", Msg: "пустое поле"}}
	}

	// TODO: При успехе верните некоторый id (можно фиксированный), который будет напечатан как OK id=<id>.
	return 1, nil
}

func GetTask(id int) (Task, error) {
	// TODO: Реализуйте команду get.
	// TODO: Если id != 1, верните ошибку, обёрнутую в *OpError (Op="GetTask"),
	// TODO: а причиной должна быть *NotFoundError с корректным ID.
	if id == 1 {
		return demoTask, nil
	}

	return Task{}, &OpError{Op: "GetTask", Err: &NotFoundError{ID: id}}
}

// ToUserError классифицирует ошибку строго через errors.As (без подстрок).
func ToUserError(err error) (code, message string, field, detail string) {
	// TODO: Классифицируйте err строго через errors.As по всей цепочке:
	// TODO: 1) ValidationError -> code="validation", message="Некорректные данные", field/detail из ValidationError.
	// TODO: 2) *NotFoundError -> code="not_found", message="Задача не найдена".
	// TODO: 3) иначе -> code="internal", message="Не удалось выполнить операцию".
	var ve ValidationError
	var nfe *NotFoundError
	switch {
	case errors.As(err, &ve):
		return "validation", "Некорректные данные", ve.Field, ve.Msg
	case errors.As(err, &nfe):
		return "not_found", "Задача не найдена", "", ""
	default:
		return "internal", "Не удалось выполнить операцию", "", ""
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(line)

	parts := strings.Fields(line)
	if len(parts) == 0 {
		printErr(&OpError{Op: "Parse", Err: errors.New("empty command")})
		return
	}

	switch parts[0] {
	case "add":
		if len(parts) != 2 {
			printErr(&OpError{Op: "Parse", Err: errors.New("bad add command")})
			return
		}

		id, err := AddTask(parts[1])
		if err != nil {
			printErr(err)
			return
		}

		fmt.Printf("OK id=%d\n", id)

	case "get":
		if len(parts) != 2 {
			printErr(&OpError{Op: "Parse", Err: errors.New("bad get command")})
			return
		}

		id, convErr := strconv.Atoi(parts[1])
		if convErr != nil {
			printErr(&OpError{Op: "Parse", Err: convErr})
			return
		}

		task, err := GetTask(id)
		if err != nil {
			printErr(err)
			return
		}

		fmt.Printf("TASK id=%d title=%s done=%v\n", task.ID, task.Title, task.Done)

	default:
		printErr(&OpError{Op: "Parse", Err: errors.New("unknown command")})
	}
}

func printErr(err error) {
	code, msg, field, detail := ToUserError(err)

	fmt.Printf("CODE=%s\n", code)
	fmt.Printf("MESSAGE=%s\n", msg)
	if code == "validation" {
		fmt.Printf("FIELD=%s DETAIL=%s\n", field, detail)
	}
	fmt.Printf("DEBUG=%s\n", err.Error())
}
