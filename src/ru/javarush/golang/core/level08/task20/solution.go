package main

import "fmt"

// Status — пользовательский тип статуса задачи (не "голый" int).
type Status int

const (
	// Enum-значения через iota.
	StatusNew Status = iota
	StatusInProgress
	StatusDone
)

// TaskFlags — пользовательский тип набора флагов (не "голый" uint).
type TaskFlags uint

const (
	// Флаги через битовые сдвиги: 1, 2, 4, ...
	FlagUrgent TaskFlags = 1 << iota
	FlagHasDeadline
	FlagBlocked
)

func statusText(s Status) string {
	// TODO: Реализуйте преобразование статуса в текст через switch:
	// new / in_progress / done, а для неизвестных значений — "unknown".
	switch s {
	case StatusNew:
		{
			return "new"
		}
	case StatusInProgress:
		{
			return "in_progress"
		}
	case StatusDone:
		{
			return "done"
		}
	default:
		return "unknown"
	}
}

func flagsText(f TaskFlags) string {
	// TODO: Реализуйте преобразование флагов в текст:
	// проверьте включённость флагов побитовыми операциями (&) в порядке:
	// urgent, затем deadline, затем blocked.
	// Соединяйте включённые флаги символом '|' без пробелов.
	// Если не включён ни один известный флаг — верните "none".
	text := ""
	if f&FlagUrgent != 0 {
		text += "urgent|"
	}
	if f&FlagHasDeadline != 0 {
		text += "deadline|"
	}
	if f&FlagBlocked != 0 {
		text += "blocked|"
	}
	if text == "" {
		return "none"
	}
	return text[:len(text)-1]
}

func main() {
	// По требованию: читаем ровно два числа через fmt.Scan(&statusCode, &flagsCode).
	var statusCode int
	var flagsCode uint
	fmt.Scan(&statusCode, &flagsCode)

	s := Status(statusCode)
	f := TaskFlags(flagsCode)

	// По требованию: ровно две строки.
	fmt.Printf("status=%s\nflags=%s\n", statusText(s), flagsText(f))
}
