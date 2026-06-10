package main

import "fmt"

func main() {
	var capacitySlots, usedSlots uint

	fmt.Scan(&capacitySlots, &usedSlots)

	if usedSlots > capacitySlots {
		fmt.Print("error\n")
		fmt.Printf("%v (%T)\n", uint(0), uint(0))
		return
	}

	var freeSlots uint
	if usedSlots <= capacitySlots {
		freeSlots = capacitySlots - usedSlots
		fmt.Printf("%v (%T)\n", freeSlots, freeSlots)
	}
}
