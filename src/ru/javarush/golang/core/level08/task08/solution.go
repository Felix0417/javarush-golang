package main

import "fmt"

// Смысловые типы, чтобы не мешать разные значения в одном int.
type TaskID int
type Status int
type Days int

const (
	StatusNew  Status = 0
	StatusDone Status = 1
)

func main() {
	var rawID, rawStatus, rawDays int
	fmt.Scan(&rawID, &rawStatus, &rawDays)

	// Явные конверсии из "сырого" int в доменные типы.
	id := TaskID(rawID)
	status := Status(rawStatus)
	days := Days(rawDays)

	if status == StatusNew {
		fmt.Printf("task #%d: %d\n", id, status)
	} else if status == StatusDone {
		fmt.Printf("task #%d: %d %v\n", id, status, days)
	} else {
		fmt.Printf("task #%d: status not implemented (%d)\n", id, rawStatus)
	}

}
