package main

import "fmt"

// Про экспортируемые идентификаторы пакета:
func isExportedName(name string) bool {
	// TODO: Реализуйте проверку экспортируемости имени по первому байту строки.
	// TODO: Имя считается экспортируемым только если первый байт в диапазоне 'A'–'Z'.
	// TODO: Пустую строку (len == 0) считайте неэкспортируемой.
	if len(name) == 0 {
		return false
	} else if name[0] >= 'A' && name[0] <= 'Z' {
		return true
	}
	return false
}

func main() {
	var name string
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Print("invalid")
		return
	}

	if isExportedName(name) {
		fmt.Print("exported")
	} else {
		fmt.Print("unexported")
	}
}
