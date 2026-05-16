package main

import "fmt"

func main() {
	var (
		title   string
		planned int
		spent   float64
		done    bool
	)

	fmt.Scan(&title, &planned, &spent, &done)

	isBig := planned >= 5

	isLong := spent >= 3.0

	needsAttention := (!done) && (isBig || isLong)

	status := title + ": IN PROGRESS"
	if done == true {
		status = title + ": DONE"
	}

	// Ровно три строки и в точном порядке.
	fmt.Println(status)
	fmt.Println("isBig=" + fmt.Sprint(isBig))
	fmt.Println("needsAttention=" + fmt.Sprint(needsAttention))
}
