package main

import "fmt"

// rectArea считает площадь прямоугольника; никакого I/O внутри быть не должно.
func rectArea(width, height int) int {
	// TODO: Реализуйте вычисление площади прямоугольника по целочисленным ширине и высоте.
	return width * height
}

// rectPerimeter считает периметр прямоугольника; никакого I/O внутри быть не должно.
func rectPerimeter(width, height int) int {
	// TODO: Реализуйте вычисление периметра прямоугольника по целочисленным ширине и высоте.
	return (width + height) * 2
}

func main() {
	var roomWidth, roomHeight int
	fmt.Scan(&roomWidth, &roomHeight)

	area := rectArea(roomWidth, roomHeight)
	perimeter := rectPerimeter(roomWidth, roomHeight)

	fmt.Printf("area=%d\n", area)
	fmt.Printf("perimeter=%d\n", perimeter)
}
