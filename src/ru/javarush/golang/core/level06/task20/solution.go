package main

import "fmt"

var appName string = "MiniApp"
var debug bool // по умолчанию false

func main() {
	var flag int
	fmt.Scan(&flag)

	var mode string

	if flag == 1 {
		debug = true
		mode = "debug"
	} else if flag == 0 {
		debug = false
		mode = "release"
	}

	fmt.Printf("%s mode=%s\n", appName, mode)
}
