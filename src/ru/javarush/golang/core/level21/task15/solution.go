package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Profile struct {
	Nick   string `json:"nick"`
	Note   string `json:"note,omitempty"`
	Age    int    `json:"age,omitempty"`
	Hidden string `json:"-"`
}

func main() {
	t := reflect.TypeOf(Profile{})

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		tag := f.Tag.Get("json")
		name, omitempty, ignored := parseJSONTag(tag)

		fmt.Printf("%s name=%s omitempty=%t ignored=%t\n", f.Name, name, omitempty, ignored)
	}
}

func parseJSONTag(tag string) (name string, omitempty bool, ignored bool) {
	// TODO: Разобрать json-тег на имя и опции (имя до первой запятой, опции после запятой).
	// TODO: Реализовать определение omitempty: true, если среди опций есть "omitempty".
	// TODO: Реализовать ignored: true, если тег равен "-", и в этом случае omitempty должно быть false.

	if tag == "" {
		return "", false, false
	}
	if tag == "-" {
		return "-", false, true
	}

	parts := strings.Split(tag, ",")
	name = parts[0]

	for _, v := range parts[1:] {
		if v == "omitempty" {
			omitempty = true
			break
		}
	}

	return name, omitempty, false
}
