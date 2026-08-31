package main

import (
	"fmt"
	"os"
)

// totalWithBonuses суммирует базу и любое количество бонусов.
// Важно: variadic-параметр (...int) должен быть последним.
func totalWithBonuses(base int, bonuses ...int) int {
	// TODO: Реализуйте суммирование base и всех бонусов.
	// TODO: Используйте цикл по bonuses (например, for/range) и прибавляйте каждый бонус к итоговой сумме.
	for _, b := range bonuses {
		base += b
	}
	return base
}

func main() {
	var base, bonusCount int
	fmt.Fscan(os.Stdin, &base, &bonusCount)

	var result int
	switch bonusCount {
	case 0:
		// TODO: Вызовите totalWithBonuses(base) без бонусов и сохраните результат в result.
		result = totalWithBonuses(base)
	case 1:
		var b1 int
		fmt.Fscan(os.Stdin, &b1)
		// TODO: Вызовите totalWithBonuses(base, b1) (бонус отдельным аргументом) и сохраните результат в result.
		result = totalWithBonuses(base, b1)
	case 2:
		var b1, b2 int
		fmt.Fscan(os.Stdin, &b1, &b2)
		// TODO: Вызовите totalWithBonuses(base, b1, b2) (оба бонуса отдельными аргументами) и сохраните результат в result.
		result = totalWithBonuses(base, b1, b2)
	}

	fmt.Println(result)
}
