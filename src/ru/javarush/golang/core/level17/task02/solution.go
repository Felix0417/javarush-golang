package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parseInt(mode, value string) (int, error) {
	n, err := strconv.Atoi(value)
	// TODO: Реализуйте два режима формирования ошибки:
	// TODO: - mode == "v": ошибка должна быть сформирована без wrapping (чтобы errors.Unwrap вернул nil)
	// TODO: - mode == "w": ошибка должна быть сформирована с wrapping (чтобы errors.Unwrap вернул исходную причину)
	if err != nil {
		switch mode {
		case "v":
			return 0, fmt.Errorf("parse %v", err)
		case "w":
			return 0, fmt.Errorf("parse %w", err)
		default:
			return 0, err
		}
	}

	return n, nil
}

func main() {
	var mode, value string
	fmt.Scan(&mode, &value)

	n, err := parseInt(mode, value)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		fmt.Printf("UNWRAP: %v\n", errors.Unwrap(err))
		return
	}

	fmt.Printf("OK: %d\n", n)
}
