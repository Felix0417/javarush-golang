package calc

import "strconv"

// Apply выполняет арифметическую операцию над a и b.
// Неизвестная операция: строго (0, false) без паники и без I/O.
func Apply(op string, a, b int) (int, bool) {
	switch op {
	case "add":
		return a + b, true
	case "sub":
		// TODO: Реализуйте операцию вычитания (a - b) и верните ok=true.
		return a - b, true
	case "mul":
		// TODO: Реализуйте операцию умножения (a * b) и верните ok=true.
		return a * b, true
	default:
		return 0, false
	}
}

// Format собирает строку вывода без fmt: только strconv.Itoa.
func Format(result int, debug bool) string {
	s := "result=" + strconv.Itoa(result)
	if debug {
		// TODO: Реализуйте debug-форматирование (другой префикс строки), не используя fmt.
		return "DEBUG " + s
	}
	return s
}
