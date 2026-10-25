package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type TodoItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Note  string `json:"note,omitempty"`
}

func ValidateJSONTags(v any) error {
	t := reflect.TypeOf(v)
	if t == nil {
		return errors.New("expected struct value, got nil")
	}

	// По условию принимаем значение структуры, не указатель.
	if t.Kind() == reflect.Ptr {
		return fmt.Errorf("expected struct value (not pointer), got %s", t)
	}
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct value, got %s", t.Kind())
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if !f.IsExported() {
			continue
		}

		tag := f.Tag.Get("json")
		if tag == "" {
			return fmt.Errorf("field %s: missing json tag", f.Name)
		}

		parts := strings.Split(tag, ",")
		name := parts[0]

		// TODO: провалидировать имя json-поля (часть тега до запятой) по правилам задачи.
		if !validJSONName(name) {
			return fmt.Errorf("field %s: invalid json name %q", f.Name, name)
		}

		// TODO: реализовать проверку опций тега (например, omitempty) по правилам задачи,
		// TODO: включая отдельные требования для полей Title и Note.
		for _, part := range parts[1:] {
			if f.Name == "Title" && part == "omitempty" {
				return fmt.Errorf("field %s: missing required json tag %q", f.Name, part)
			}
			if f.Name == "Note" && part != "omitempty" {
				return fmt.Errorf("field %s: missing required json tag %q", f.Name, part)
			}
		}
	}

	return nil
}

func validJSONName(s string) bool {
	// TODO: реализовать строгую проверку имени json-поля по правилам задачи.
	return regexp.MustCompile("^[a-z-]+$").MatchString(s)
}

func main() {
	var item TodoItem

	if err := ValidateJSONTags(item); err == nil {
		fmt.Println("OK")
	} else {
		fmt.Println("ERR:", err)
	}
}
