package main

import (
	"fmt"
	"os"
	"strconv"
)

func parsePrice(priceStr string) (int, error) {
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		// TODO: Верните ошибку парсинга цены в требуемом формате, сохранив исходную строку.
		return 0, fmt.Errorf("parse price %q: %v", priceStr, err)
	}
	return price, nil
}

func parsePercent(percentStr string) (int, error) {
	percent, err := strconv.Atoi(percentStr)
	if err != nil {
		// TODO: Верните ошибку парсинга процента в требуемом формате, сохранив исходную строку.
		return 0, fmt.Errorf("parse percent %q: %v", percentStr, err)
	}

	// TODO: Проверьте диапазон процента (включительно) и верните ошибку требуемого формата, если он нарушен.
	if percent < 0 || percent > 100 {
		return 0, fmt.Errorf("percent must be between 0 and 100, got %d", percent)
	}
	return percent, nil
}

func calculate(priceStr, percentStr string) (int, error) {
	price, err := parsePrice(priceStr)
	if err != nil {
		return 0, err
	}

	percent, err := parsePercent(percentStr)
	if err != nil {
		return 0, err
	}

	discounted := price - price*percent/100
	return discounted, nil
}

func main() {
	var priceStr, percentStr string
	fmt.Fscan(os.Stdin, &priceStr, &percentStr)

	discounted, err := calculate(priceStr, percentStr)
	if err != nil {
		// TODO: Добавьте верхний контекст ошибки для сценария расчёта скидки (через fmt.Errorf), как требуется в условии.
		err = fmt.Errorf("calculate: %v", err)
		fmt.Println(err)
		return
	}

	fmt.Printf("discounted=%d", discounted)
}
