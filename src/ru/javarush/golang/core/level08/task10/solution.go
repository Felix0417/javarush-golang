package main

import "fmt"

type Command int
type Priority int

const (
	// TODO: Описать перечисление команд через iota в отдельном const-блоке:
	// TODO: CmdStart=0, CmdStop=1, CmdRestart=2 (тип Command).
	CmdStart Command = iota
	CmdStop
	CmdRestart
)

const (
	// TODO: Описать перечисление приоритетов через iota во втором (независимом) const-блоке:
	// TODO: PrLow=0, PrHigh=1 (тип Priority).
	PrLow Priority = iota
	PrHigh
)

func main() {
	fmt.Println(CmdStart, CmdStop, CmdRestart)
	fmt.Println(PrLow, PrHigh)
}
