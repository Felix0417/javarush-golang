package main

import "fmt"

func main() {
	// Мини-словарь "английское слово -> перевод" (обязательно литералом, без make).
	dictionary := map[string]string{
		"hello":  "привет",  // TODO: Заполните перевод для слова "hello".
		"bye":    "пока",    // TODO: Заполните перевод для слова "bye".
		"thanks": "спасибо", // TODO: Заполните перевод для слова "thanks".
	}

	var inputWord string
	// TODO: Считайте одно английское слово из stdin в переменную inputWord с помощью fmt.Scan или fmt.Fscan.
	fmt.Scan(&inputWord)
	// По условию слово всегда есть в словаре — ok-проверка запрещена/не нужна.
	fmt.Println(dictionary[inputWord])
	fmt.Println(len(dictionary))
}
