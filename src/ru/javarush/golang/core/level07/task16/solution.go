package main

import "fmt"

func main() {
	var firstByteValue, secondByteValue int
	fmt.Scan(&firstByteValue, &secondByteValue)

	outOfRange := (firstByteValue < 0 || firstByteValue > 255) || (secondByteValue < 0 || secondByteValue > 255)

	if outOfRange {
		fmt.Print("OUT")
		return
	}

	var firstByte, secondByte = uint8(firstByteValue), uint8(secondByteValue)

	var wrappedSum uint8 = firstByte + secondByte

	var exactSum int = int(firstByte) + int(secondByte)

	fmt.Printf("%d %d", wrappedSum, exactSum)
}
