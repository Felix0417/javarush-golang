package main

import "fmt"

// NotifyFlag — набор флагов уведомлений, хранится как битовая маска.
type NotifyFlag uint

const (
	// TODO: объявите флаги FlagEmail, FlagSMS, FlagPush в одном const-блоке как битовые флаги через iota (паттерн 1 << iota).
	FlagEmail NotifyFlag = 1 << iota
	FlagSMS
	FlagPush
)

func main() {
	var raw uint
	fmt.Scan(&raw)

	mask := NotifyFlag(raw)

	if mask == 0 {
		fmt.Println("no flags")
		return
	}

	// TODO: реализуйте корректную проверку включённости флагов как битовой маски (поддержите комбинации флагов).
	// TODO: обеспечьте вывод включённых каналов строго в порядке: email, затем sms, затем push (каждый с новой строки).
	// Временная упрощённая логика: распознаёт только ровно один установленный флаг.
	if mask&FlagEmail != 0 {
		fmt.Println("email")
	}
	if mask&FlagSMS != 0 {
		fmt.Println("sms")
	}
	if mask&FlagPush != 0 {
		fmt.Println("push")
	}
}
