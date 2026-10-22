package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseQuantity(s string) (int, error) {
	// TODO: Распарсить строку в целое число через strconv.Atoi; при ошибке вернуть (0, error) с корректным текстом ошибки
	s = strings.ToLower(s)
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parsing %q: %v", s, err)
	}

	// TODO: Если парсинг успешен — проверить диапазон 1..1000 включительно; при нарушении вернуть (0, error) с корректным текстом ошибки
	if num <= 0 || num > 1000 {
		return 0, fmt.Errorf("value out of range between 1 and 1000, got %d", num)
	}
	return num, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.TrimSpace(scanner.Text())

	quantity, err := ParseQuantity(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("OK %d\n", quantity)
}
