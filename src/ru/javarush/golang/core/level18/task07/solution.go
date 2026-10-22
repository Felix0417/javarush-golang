package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	errWriteFailed = errors.New("write failed")
	errCloseFailed = errors.New("close failed")
)

type FakeFile struct {
	closeFails bool
}

func (f FakeFile) WriteLine(s string) error {
	// TODO: Реализуйте логику ошибки записи: при s == "BAD" вернуть errWriteFailed, иначе nil.
	if s == "BAD" {
		return errWriteFailed
	}
	return nil
}

func (f FakeFile) Close() error {
	// TODO: Реализуйте логику ошибки закрытия: при f.closeFails == true вернуть errCloseFailed, иначе nil.
	if f.closeFails {
		return errCloseFailed
	}
	return nil
}

func SaveLine(mode string, line string) (err error) {
	f := FakeFile{closeFails: false}
	// TODO: Настройте f.closeFails в зависимости от mode (режим "closefail").
	if mode == "closefail" {
		f.closeFails = true
	}

	// TODO: Оформите defer через анонимную функцию.
	// TODO: В defer вызовите Close() и обработайте ошибку закрытия так, чтобы она не затирала ошибку записи.

	defer func() {
		cerr := f.Close()
		if cerr != nil && err == nil {
			err = cerr
		}
	}()

	err = f.WriteLine(line)
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var mode, line string
	fmt.Fscan(in, &mode, &line)

	err := SaveLine(mode, line)
	fmt.Printf("err: %v\n", err)
}
