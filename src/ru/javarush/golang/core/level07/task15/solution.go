package main

import "fmt"

func main() {
	var incomingValue int64
	fmt.Scan(&incomingValue)

	fitsInt8 := (incomingValue >= -128) && (incomingValue <= 127)

	if fitsInt8 {
		fmt.Println(int8(incomingValue))
	} else {
		fmt.Println("OVERFLOW")
	}
}
