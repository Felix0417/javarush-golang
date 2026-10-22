package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseOptionalInt(s string) (*int, error) {
	if s == "-" {
		return nil, nil
	}

	// TODO: Реализуйте парсинг целого числа из строки (через strconv.Atoi).
	// TODO: При успехе верните указатель на число и nil-ошибку.
	// TODO: При ошибке парсинга верните (nil, error) с текстом строго: bad number "<s>".
	tmp, err := strconv.Atoi(s)
	if err != nil {
		return nil, fmt.Errorf("bad number %q", s)
	}
	return &tmp, nil
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	s := in.Text()

	v, err := parseOptionalInt(s)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if v == nil {
		fmt.Println("nil")
		return
	}

	fmt.Println(*v)
}
