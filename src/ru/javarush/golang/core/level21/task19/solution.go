package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type DeliveryState int

const (
	StateCreated DeliveryState = iota + 1
	StateSent
	StateDelivered
)

type Delivery struct {
	ID      int
	Address string
	State   DeliveryState
}

func ValidateDelivery(d Delivery) error {
	// TODO: Проверьте инварианты модели Delivery (ID, Address после TrimSpace, допустимые значения State).
	// TODO: Для неизвестного состояния верните ошибку с текстом "unknown state".

	if d.ID <= 0 {
		return errors.New("invalid id")
	}

	if strings.TrimSpace(d.Address) == "" {
		return errors.New("empty address")
	}

	// TODO: Реализуйте полную проверку допустимых значений State.
	if d.State < StateCreated || d.State > StateDelivered {
		return errors.New("unknown state")
	}

	return nil
}

func SetState(d *Delivery, newState DeliveryState) error {
	// TODO: При nil нужно вернуть ошибку с текстом "task is nil" и не выполнять дальнейшую логику.
	if d == nil {
		return errors.New("task is nil")
	}

	// TODO: Провалидируйте текущую модель перед попыткой перехода.
	if err := ValidateDelivery(*d); err != nil {
		return err
	}

	// TODO: Запретите переход "назад" (newState меньше текущего d.State) и верните "invalid state transition".
	if newState < d.State {
		return errors.New("invalid state transition")
	}

	// Обновляем состояние (пока без проверки корректности перехода).
	d.State = newState

	// TODO: После обновления обязателен вызов ValidateDelivery для всей модели и возврат ошибки при невалидности.
	return ValidateDelivery(*d)
}

func main() {
	var idInput int
	var addressInput string
	var oldStateInput int
	var newStateInput int

	fmt.Fscan(os.Stdin, &idInput, &addressInput, &oldStateInput, &newStateInput)

	d := Delivery{
		ID:      idInput,
		Address: addressInput,
		State:   DeliveryState(oldStateInput),
	}

	// Валидация на границе: сразу после ввода.
	if err := ValidateDelivery(d); err != nil {
		fmt.Printf("INVALID: %s", err.Error())
		return
	}

	if err := SetState(&d, DeliveryState(newStateInput)); err != nil {
		fmt.Printf("INVALID: %s", err.Error())
		return
	}

	fmt.Printf("UPDATED: %d", d.State)
}
