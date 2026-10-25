package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Visibility int

const (
	VisPrivate Visibility = iota + 1
	VisPublic
)

type Note struct {
	ID         int
	Text       string
	Visibility Visibility
}

func ValidateNote(n Note) error {
	// TODO: Реализуйте проверку инвариантов Note:
	// TODO: 1) ID должен быть > 0
	// TODO: 2) Text не должен быть пустым
	// TODO: 3) Visibility должен быть одним из допустимых значений (VisPrivate/VisPublic)
	// TODO: При нарушении каждого инварианта верните ошибку с точным текстом из условия.
	if n.ID < 1 {
		return errors.New("id must be positive")
	} else if n.Text == "" {
		return errors.New("text is empty")
	} else if n.Visibility < VisPrivate || n.Visibility > VisPublic {
		return errors.New("unknown visibility")
	}

	return nil
}

func NewNote(id int, text string, vis Visibility) (Note, error) {
	// TODO: Реализуйте единую точку сборки Note:
	// TODO: 1) Нормализуйте text через strings.TrimSpace
	// TODO: 2) Соберите структуру Note
	// TODO: 3) Вызовите ValidateNote и при ошибке верните Note{} и error

	text = strings.TrimSpace(text)

	n := Note{
		ID:         id,
		Text:       text,
		Visibility: vis,
	}

	if err := ValidateNote(n); err != nil {
		return Note{}, err
	}
	return n, nil
}

func main() {
	r := bufio.NewReader(os.Stdin)

	idLine, _ := r.ReadString('\n')
	textLine, _ := r.ReadString('\n')
	visibilityLine, _ := r.ReadString('\n')

	id, err := strconv.Atoi(strings.TrimSpace(idLine))
	if err != nil {
		fmt.Printf("INVALID INPUT: %v", err)
		return
	}

	visInt, err := strconv.Atoi(strings.TrimSpace(visibilityLine))
	if err != nil {
		fmt.Printf("INVALID INPUT: %v", err)
		return
	}
	vis := Visibility(visInt)

	n, err := NewNote(id, textLine, vis)
	if err != nil {
		fmt.Printf("INVALID NOTE: %v", err)
		return
	}

	fmt.Printf("CREATED: %+v", n)
}
