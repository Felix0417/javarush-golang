package main

import (
	"fmt"
)

func main() {
	if err := safeRun(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func safeRun() (err error) {
	// TODO: Реализуйте перехват паники на границе приложения через defer+recover и преобразование в error нужного формата.
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("unexpected panic: %v", e)
		}
	}()
	return run()
}

func run() error {
	var userToken string
	if _, scanErr := fmt.Scan(&userToken); scanErr != nil {
		return scanErr
	}

	// TODO: Реализуйте поведение: если введено слово "panic", нужно вызвать panic("panic requested").
	if userToken == "panic" {
		panic("panic requested")
	}

	fmt.Printf("ok: %s\n", userToken)
	return nil
}
