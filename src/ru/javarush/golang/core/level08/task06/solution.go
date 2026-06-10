package main

import "fmt"

// Speed — смысловой тип скорости, чтобы не путать с другими int.
type Speed int

// Step — именованный шаг скорости.
// TODO: Сделайте Step константой типа Speed и задайте шаг равным 1.
const Step Speed = 1

func main() {
	var rawSpeed int
	fmt.Scan(&rawSpeed)

	// TODO: Явно преобразуйте rawSpeed в Speed и сохраните в s.
	var s Speed = Speed(rawSpeed)

	// TODO: Посчитайте next как s + Step (в типе Speed).
	var next Speed = s + Step

	fmt.Printf("raw=%d type=%T\n", rawSpeed, rawSpeed)
	fmt.Printf("s=%d type=%T\n", s, s)
	fmt.Printf("next=%d\n", next)
}
