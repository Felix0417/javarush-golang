package main

import "fmt"

func main() {
	var login string
	var code int

	fmt.Scan(&login, &code)

	var isAdmin bool = login == "admin"

	var isCodeOK bool = code == 4321

	var canEnter bool = isAdmin && isCodeOK

	fmt.Printf("canEnter=%v\n", canEnter)

	if canEnter {
		fmt.Println("ACCESS GRANTED")
	} else {
		fmt.Println("ACCESS DENIED")
	}
}
