package main

import "fmt"

func main() {
	var levelsCount int
	fmt.Scan(&levelsCount)

	if levelsCount < 0 {
		fmt.Print("error")
		return
	}

	var totalRewards int
	for level := 1; level <= levelsCount; level++ {
		totalRewards += level
	}

	fmt.Printf("%d (%T)", totalRewards, totalRewards)
}
